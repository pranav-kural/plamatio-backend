package categories

import (
	"context"

	db "encore.app/categories/db"
	models "encore.app/categories/models"
	"encore.dev/storage/sqldb"
)

// ProductDB instance
var ProductsDB = sqldb.Named("products")

// CategoriesTable instance
var CategoriesTable = &db.CategoriesTable{DB: ProductsDB}

// GET: /categories/get/:id
// Retrieves the category from the database with the given ID.
//encore:api public method=GET path=/categories/get/:id
func GetCategory(ctx context.Context, id int) (*models.Category, error) {
	// Retrieve the product from the database.
	r, err := CategoriesTable.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// Return the product.
	return r, err
}

// GET: /categories/all
// Retrieves all categories from the database.
//encore:api public method=GET path=/categories/all
func GetCategories(ctx context.Context) (*models.Categories, error) {
	// Retrieve all products from the database.
	r, err := CategoriesTable.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	// Return the products.
	return r, err
}