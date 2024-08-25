package products

import (
	"context"

	models "encore.app/products/models"
	"encore.dev/storage/sqldb"
)

// Inserts a product into the database.
func Insert(ctx context.Context, db *sqldb.Database, name string, imageUrl string, price int) error {
	_, err := db.Exec(ctx, `
		INSERT INTO products (name, imageUrl, price)
		VALUES ($1, $2, $3)
	`, name, imageUrl, price)
	return err
}

// Retrieves a product from the database.
func Get(ctx context.Context, db *sqldb.Database, id int) (*models.Product, error) {
	p := &models.Product{ID: id}
	err := db.QueryRow(ctx, `
		SELECT name, imageUrl, price FROM products
		WHERE id = $1
	`, id).Scan(&p.Name, &p.ImageURL, &p.Price)
	return p, err
}

// Deletes a product from the database.
func Delete(ctx context.Context, db *sqldb.Database, id int) error {
	_, err := db.Exec(ctx, `
		DELETE FROM products
		WHERE id = $1
	`, id)
	return err
}

// Updates a product in the database.
func Update(ctx context.Context, db *sqldb.Database, id int, name string, imageUrl string, price int) error {
	_, err := db.Exec(ctx, `
		UPDATE products
		SET name = $1, imageUrl = $2, price = $3
		WHERE id = $4
	`, name, imageUrl, price, id)
	return err
}

// Retrieves all products from the database.
func GetAll(ctx context.Context, db *sqldb.Database) ([]*models.Product, error) {
	rows, err := db.Query(ctx, `
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