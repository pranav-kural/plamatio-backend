package products

import (
	"context"

	db "encore.app/products/db"
	models "encore.app/products/models"
)

// PDB is the ProductsDB instance.
var CategoriesTable = &db.CategoriesTable{DB: ProductsDB}

// GET: /products/get/:id
// Retrieves the product from the database with the given ID.
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

// GET: /products/all
// Retrieves all products from the database.
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