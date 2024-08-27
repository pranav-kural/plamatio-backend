package products

import (
	"context"

	categoriesModels "encore.app/categories/models"
	db "encore.app/products/db"
	models "encore.app/products/models"
	"encore.dev/storage/sqldb"
)

// Create a new database instance for the products database.
var ProductsDB = sqldb.NewDatabase("products", sqldb.DatabaseConfig{
		Migrations: "./migrations",
	})

// ProductsTB is the products table instance.
var ProductsTB = &db.ProductsTB{DB: ProductsDB}

// GET: /products/get/:id
// Retrieves the product from the database with the given ID.
//encore:api public method=GET path=/products/get/:id
func Get(ctx context.Context, id int) (*models.Product, error) {
	// Retrieve the product from the database.
	r, err := ProductsTB.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// Return the product.
	return r, err
}

// POST: /products
// Inserts a product into the database.
//encore:api private method=POST path=/products/add
func Insert(ctx context.Context, p *models.ProductRequestParams) (*models.Product, error) {
	// Insert the product into the database.
	if id, err := ProductsTB.Insert(ctx, p); err != nil {
		return nil, err
	} else {
		// Return the product.
		return &models.Product{ID: id, Name: p.Name, Description: p.Description, CategoryId: categoriesModels.ProductCategory(p.CategoryId), ImageURL: p.ImageURL, Price: p.Price}, nil
	}
}

// DELETE: /products/delete/:id
// Deletes the product from the database with the given ID.
//encore:api private method=DELETE path=/products/delete/:id
func Delete(ctx context.Context, id int) error {
	// Delete the product from the database.
	if err := ProductsTB.Delete(ctx, id); err != nil {
		return err
	}
	// Return nil if successful.
	return nil
}

// PUT: /products/update/:id
// Updates the product in the database with the given ID.
//encore:api private method=PUT path=/products/update/:id
func Update(ctx context.Context, id int, p *models.ProductRequestParams) (*models.Product, error) {
	// Update the product in the database.
	if err := ProductsTB.Update(ctx, id, p); err != nil {
		return nil, err
	}
	// Return the updated product.
	return &models.Product{ID: id, Name: p.Name, Description: p.Description, CategoryId: categoriesModels.ProductCategory(p.CategoryId), ImageURL: p.ImageURL, Price: p.Price}, nil
}

// GET: /products/all
// Retrieves all products from the database.
//encore:api public method=GET path=/products/all
func GetAll(ctx context.Context) (*models.Products, error) {
	// Retrieve all products from the database.
	r, err := ProductsTB.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	// Return the products.
	return r, err
}

// GET: /products/category/:id
// Retrieves all products from the database by category.
//encore:api public method=GET path=/products/category/:id
func GetByCategory(ctx context.Context, id int) (*models.Products, error) {
	// Retrieve all products from the database by category.
	r, err := ProductsTB.GetByCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	// Return the products.
	return r, err
}

// GET: /products/hero-products
// Retrieves all hero products from the database.
//encore:api public method=GET path=/products/hero-products
func GetHeroProducts(ctx context.Context) (*models.Products, error) {
	// Retrieve all hero products from the database.
	r, err := ProductsTB.GetHeroProducts(ctx)
	if err != nil {
		return nil, err
	}
	// Return the products.
	return r, err
}

// GET: /products/search/:query
// Retrieves all products from the database by search query.
//encore:api public method=GET path=/products/search/:query
func Search(ctx context.Context, query string) (*models.Products, error) {
	// Retrieve all products from the database by search query.
	r, err := ProductsTB.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	// Return the products.
	return r, err
}