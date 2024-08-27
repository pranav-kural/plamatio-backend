package products

import (
	"context"

	models "encore.app/products/models"
	utils "encore.app/products/utils"
	"encore.dev/storage/sqldb"
)

type ProductsDB struct {
	DB *sqldb.Database
}

const (
    SQL_INSERT_PRODUCT = `
        INSERT INTO products (name, description, category_id, image_url, price)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
    SQL_BULK_INSERT_PRODUCTS = `
        INSERT INTO products (name, description, category_id, image_url, price)
        VALUES ($1, $2, $3, $4, $5)
    `
    SQL_DELETE_PRODUCT = `
        DELETE FROM products
        WHERE id = $1
    `
    SQL_UPDATE_PRODUCT = `
        UPDATE products
        SET name = $1, description = $2, category_id = $3, image_url = $4, price = $5
        WHERE id = $6
    `
		SQL_GET_PRODUCT = `
        SELECT name, description, category_id, image_url, price FROM products
        WHERE id = $1
    `
    SQL_GET_ALL_PRODUCTS = `
        SELECT id, name, description, category_id, image_url, price FROM products
    `
		SQL_GET_PRODUCTS_BY_CATEGORY = `
				SELECT id, name, description, category_id, image_url, price FROM products
				WHERE category_id = $1
		`
		SQL_GET_HERO_PRODUCTS = `
				SELECT p.id, p.name, p.description, p.category_id, p.image_url, p.price
				FROM products p
				INNER JOIN hero_products hp ON p.id = hp.product_id
		`
)

// Inserts a product into the database.
// Inserts a product into the database and returns the id of the newly added record.
func (pdb *ProductsDB) Insert(ctx context.Context, p *models.ProductRequestParams) (int, error) {
    var id int
    err := pdb.DB.QueryRow(ctx, SQL_INSERT_PRODUCT, p.Name, p.Description, p.CategoryId, p.ImageURL, p.Price).Scan(&id)
		
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
		if _, err := stmt.Exec(p.Name, p.Description, p.CategoryId, p.ImageURL, p.Price); err != nil {
			return err
		}
	}
  return nil
}

// Deletes a product from the database.
func (pdb *ProductsDB) Delete(ctx context.Context, id int) error {
	_, err := pdb.DB.Exec(ctx, SQL_DELETE_PRODUCT, id)
	return err
}

// Updates a product in the database.
func (pdb *ProductsDB) Update(ctx context.Context, id int, p *models.ProductRequestParams) error {
	_, err := pdb.DB.Exec(ctx, SQL_UPDATE_PRODUCT, p.Name, p.Description, p.CategoryId, p.ImageURL, p.Price, id)
	return err
}

// Retrieves a product from the database.
func (pdb *ProductsDB) Get(ctx context.Context, id int) (*models.Product, error) {
	p := &models.Product{ID: id}
	err := pdb.DB.QueryRow(ctx, SQL_GET_PRODUCT, id).Scan(&p.Name, &p.Description, &p.CategoryId, &p.ImageURL, &p.Price)
	return p, err
}

// Retrieves all products from the database.
func (pdb *ProductsDB) GetAll(ctx context.Context) (*models.Products, error) {
	rows, err := pdb.DB.Query(ctx, SQL_GET_ALL_PRODUCTS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CategoryId, &p.ImageURL, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &models.Products{Data: products}, nil
}

// Retrieves all products from the database by category.
func (pdb *ProductsDB) GetByCategory(ctx context.Context, id int) (*models.Products, error) {
	rows, err := pdb.DB.Query(ctx, SQL_GET_PRODUCTS_BY_CATEGORY, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CategoryId, &p.ImageURL, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &models.Products{Data: products}, nil
}

/*
CREATE TABLE hero_products (
    category_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    PRIMARY KEY (category_id, product_id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
*/

// Retrieve hero products from the database.
func (pdb *ProductsDB) GetHeroProducts(ctx context.Context) (*models.Products, error) {
	// Query the database for hero products.
	rows, err := pdb.DB.Query(ctx, SQL_GET_HERO_PRODUCTS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and create a product for each row.
	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CategoryId, &p.ImageURL, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &models.Products{Data: products}, nil
}

// Retrieve products based on search query.
func (pdb *ProductsDB) Search(ctx context.Context, query string) (*models.Products, error) {
	// Query the database for products based on search query.
	rows, err := pdb.DB.Query(ctx, "SELECT id, name, description, category_id, image_url FROM products WHERE name ILIKE $1", "%"+query+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and create a product for each row.
	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CategoryId, &p.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return &models.Products{Data: products}, nil
}