package products

import (
	"context"

	db "encore.app/products/db"
	models "encore.app/products/models"
)

// GET: /products/:id
// Retrieves the product from the database with the given ID.
//encore:api public method=GET path=/products/:id
func Get(ctx context.Context, id int) (*models.Product, error) {
	// Get the products database instance.
	pdb, err := db.GetProductsDB()
	if err != nil {
		return nil, err
	}
	// Retrieve the product from the database.
	r, err := pdb.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// Return the product.
	return r, err
}

// POST: /products
// Inserts a product into the database.
//encore:api public method=POST path=/products
func Insert(ctx context.Context, name string, imageUrl string, price int) (*models.Product, error) {
	// Get the products database instance.
	pdb, err := db.GetProductsDB()
	if err != nil {
		return nil, err
	}
	// Insert the product into the database.
	if id, err := pdb.Insert(ctx, name, imageUrl, price); err != nil {
		return nil, err
	} else {
		// Return the product.
		return &models.Product{ID: id, Name: name, ImageURL: imageUrl, Price: price}, nil
	}
}

// DELETE: /products/:id
// Deletes the product from the database with the given ID.
//encore:api public method=DELETE path=/products/:id
func Delete(ctx context.Context, id int) error {
	// Get the products database instance.
	pdb, err := db.GetProductsDB()
	if err != nil {
		return err
	}
	// Delete the product from the database.
	if err := pdb.Delete(ctx, id); err != nil {
		return err
	}
	// Return nil if successful.
	return nil
}

// PUT: /products/:id
// Updates the product in the database with the given ID.
//encore:api public method=PUT path=/products/:id
func Update(ctx context.Context, id int, name string, imageUrl string, price int) (*models.Product, error) {
	// Get the products database instance.
	pdb, err := db.GetProductsDB()
	if err != nil {
		return nil, err
	}
	// Update the product in the database.
	if err := pdb.Update(ctx, id, name, imageUrl, price); err != nil {
		return nil, err
	}
	// Return the updated product.
	return &models.Product{ID: id, Name: name, ImageURL: imageUrl, Price: price}, nil
}

// GET: /products/all
// Retrieves all products from the database.
//encore:api public method=GET path=/products/all
func GetAll(ctx context.Context) ([]*models.Product, error) {
	// Get the products database instance.
	pdb, err := db.GetProductsDB()
	if err != nil {
		return nil, err
	}
	// Retrieve all products from the database.
	r, err := pdb.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	// Return the products.
	return r, err
}

