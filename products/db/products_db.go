package products

import (
	"context"

	models "encore.app/products/models"
	"encore.dev/storage/sqldb"
)

type ProductsDB struct {
	db *sqldb.Database
}

func GetProductsDB() (*ProductsDB, error) {
	const PRODUCTS_DB_NAME = "products"
	db := sqldb.Named(PRODUCTS_DB_NAME)
	return &ProductsDB{db}, nil
}


func CreateProductsTable(ctx context.Context) (*sqldb.Database, error) {
	var db = sqldb.NewDatabase("products", sqldb.DatabaseConfig{
		Migrations: "./migrations",
	})
	return db, nil
}

// Inserts a product into the database.
// Inserts a product into the database and returns the id of the newly added record.
func (pdb *ProductsDB) Insert(ctx context.Context, name string, imageUrl string, price int) (int, error) {
    var id int
    err := pdb.db.QueryRow(ctx, `
        INSERT INTO products (name, imageUrl, price)
        VALUES ($1, $2, $3)
        RETURNING id
    `, name, imageUrl, price).Scan(&id)

    return id, err
}

// Retrieves a product from the database.
func (pdb *ProductsDB) Get(ctx context.Context, id int) (*models.Product, error) {
	p := &models.Product{ID: id}
	err := pdb.db.QueryRow(ctx, `
		SELECT name, imageUrl, price FROM products
		WHERE id = $1
	`, id).Scan(&p.Name, &p.ImageURL, &p.Price)
	return p, err
}

// Deletes a product from the database.
func (pdb *ProductsDB) Delete(ctx context.Context, id int) error {
	_, err := pdb.db.Exec(ctx, `
		DELETE FROM products
		WHERE id = $1
	`, id)
	return err
}

// Updates a product in the database.
func (pdb *ProductsDB) Update(ctx context.Context, id int, name string, imageUrl string, price int) error {
	_, err := pdb.db.Exec(ctx, `
		UPDATE products
		SET name = $1, imageUrl = $2, price = $3
		WHERE id = $4
	`, name, imageUrl, price, id)
	return err
}

// Retrieves all products from the database.
func (pdb *ProductsDB) GetAll(ctx context.Context) ([]*models.Product, error) {
	rows, err := pdb.db.Query(ctx, `
		SELECT id, name, imageUrl, price FROM products
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.ImageURL, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}