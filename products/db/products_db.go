package products

import (
	"context"

	models "encore.app/products/models"
	utils "encore.app/products/utils"
	rLog "encore.dev/rlog"
	"encore.dev/storage/sqldb"
)

type ProductsDB struct {
	DB *sqldb.Database
}

const (
    SQL_INSERT_PRODUCT = `
        INSERT INTO products (name, description, category, imageurl, price)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
    SQL_BULK_INSERT_PRODUCTS = `
        INSERT INTO products (name, description, category, imageurl, price)
        VALUES ($1, $2, $3, $4, $5)
    `
    SQL_GET_PRODUCT = `
        SELECT name, description, category, imageurl, price FROM products
        WHERE id = $1
    `
    SQL_DELETE_PRODUCT = `
        DELETE FROM products
        WHERE id = $1
    `
    SQL_UPDATE_PRODUCT = `
        UPDATE products
        SET name = $1, description = $2, category = $3, imageurl = $4, price = $5
        WHERE id = $6
    `
    SQL_GET_ALL_PRODUCTS = `
        SELECT id, name, description, category, imageurl, price FROM products
    `
)

// Inserts a product into the database.
// Inserts a product into the database and returns the id of the newly added record.
func (pdb *ProductsDB) Insert(ctx context.Context, p *models.ProductRequestParams) (int, error) {
    var id int
    err := pdb.DB.QueryRow(ctx, SQL_INSERT_PRODUCT, p.Name, p.Description, p.Category, p.ImageURL, p.Price).Scan(&id)
		
    return id, err
}

// Bulk inserts products into the database.
func (pdb *ProductsDB) BulkInsert(ctx context.Context, products []*models.ProductRequestParams) error {
	tx, err := pdb.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := pdb.DB.Stdlib().Prepare(SQL_BULK_INSERT_PRODUCTS)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range products {
		// validate data
		err := utils.ValidateProductRequestParams(p)
		if err != nil {
			return err
		}
		if _, err := stmt.Exec(p.Name, p.Description, p.Category, p.ImageURL, p.Price); err != nil {
			return err
		}
	}
  return nil
}

// Retrieves a product from the database.
func (pdb *ProductsDB) Get(ctx context.Context, id int) (*models.Product, error) {
	p := &models.Product{ID: id}
	err := pdb.DB.QueryRow(ctx, SQL_GET_PRODUCT, id).Scan(&p.Name, &p.Description, &p.Category, &p.ImageURL, &p.Price)
	return p, err
}

// Deletes a product from the database.
func (pdb *ProductsDB) Delete(ctx context.Context, id int) error {
	_, err := pdb.DB.Exec(ctx, SQL_DELETE_PRODUCT, id)
	return err
}

// Updates a product in the database.
func (pdb *ProductsDB) Update(ctx context.Context, id int, p *models.ProductRequestParams) error {
	_, err := pdb.DB.Exec(ctx, SQL_UPDATE_PRODUCT, p.Name, p.Description, p.Category, p.ImageURL, p.Price, id)
	return err
}

// Retrieves all products from the database.
func (pdb *ProductsDB) GetAll(ctx context.Context) (*models.Products, error) {
	rLog.Info("Retrieving all products from the database")
	rows, err := pdb.DB.Query(ctx, SQL_GET_ALL_PRODUCTS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Category, &p.ImageURL, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &models.Products{Data: products}, nil
}