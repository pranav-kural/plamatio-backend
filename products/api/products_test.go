package products

import (
	"context"
	"testing"

	models "encore.app/products/models"
)

// Run tests using `encore test`, which compiles the Encore app and then runs `go test`.
// It supports all the same flags that the `go test` command does.
// You automatically get tracing for tests in the local dev dash: http://localhost:9400
// Learn more: https://encore.dev/docs/develop/testing

// TestProductInsertUpdateGetDelete - tests inserting, updating, retrieving, and deleting a product from the database.
func TestProductInsertUpdateGetDelete(t *testing.T) {
	testProductParams := &models.ProductRequestParams{
		Name:     "Test Product",
		Description: "Test Description",
		ImageURL: "https://example.com/image.jpg",
		CategoryId: 1,
		Price:    1000,
	}
	addedProduct, err := Insert(context.Background(), testProductParams)
	if err != nil {
		t.Fatal(err)
	}
	retrievedProduct, err := Get(context.Background(), addedProduct.ID)
	if err != nil {
		t.Fatal(err)
	}
	if retrievedProduct.Name != testProductParams.Name || retrievedProduct.ImageURL != testProductParams.ImageURL || retrievedProduct.Price != testProductParams.Price {
		t.Errorf("got %v, want %v", retrievedProduct, testProductParams)
	}
	ProductRequestParams := &models.ProductRequestParams{
		Name: "Updated Product",
		Description: "Updated Description",
		ImageURL: "https://example.com/updated.jpg",
		CategoryId: 2,
		Price: 2000,
	}
	updatedProduct, err := Update(context.Background(), addedProduct.ID, ProductRequestParams)
	if err != nil {
		t.Fatal(err)
	}
	if updatedProduct.Name != ProductRequestParams.Name || updatedProduct.ImageURL != ProductRequestParams.ImageURL || updatedProduct.Price != ProductRequestParams.Price {
		t.Errorf("got %v, want %v", updatedProduct, ProductRequestParams)
	}
	err = Delete(context.Background(), addedProduct.ID)
	if err != nil {
		t.Fatal(err)
	}
}

// TestProductsGetAll - tests retrieving all products from the database.
func TestProductsGetAll(t *testing.T) {
	// number of products initially loaded to the database through migrations
	const NUM_PRODUCTS = 35
	retrievedProducts, err := GetAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(retrievedProducts.Data) < NUM_PRODUCTS {
		t.Errorf("got %v, want %v", len(retrievedProducts.Data), NUM_PRODUCTS)
	}
}