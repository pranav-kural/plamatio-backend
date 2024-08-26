package url

import (
	"context"
	"testing"

	api "encore.app/products/api"
	models "encore.app/products/models"
)

// Run tests using `encore test`, which compiles the Encore app and then runs `go test`.
// It supports all the same flags that the `go test` command does.
// You automatically get tracing for tests in the local dev dash: http://localhost:9400
// Learn more: https://encore.dev/docs/develop/testing

// TestProductsInsert - tests inserting a product into the database.
func TestProductsInsert(t *testing.T) {
	testProductParams := &models.ProductRequestParams{
		Name:     "Test Product",
		ImageURL: "https://example.com/image.jpg",
		Price:    1000,
	}
	addedProduct, err := api.Insert(context.Background(), testProductParams)
	if err != nil {
		t.Fatal(err)
	}
	if addedProduct.Name != testProductParams.Name || addedProduct.ImageURL != testProductParams.ImageURL || addedProduct.Price != testProductParams.Price {
		t.Errorf("got %v, want %v", addedProduct, testProductParams)
	}
}

// TestProductsBulkInsert - tests bulk inserting products into the database.
func TestProductsBulkInsert(t *testing.T) {
	testProducts := []*models.ProductRequestParams{
		{Name: "Test Product 1", ImageURL: "https://example.com/image1.jpg", Price: 1000},
		{Name: "Test Product 2", ImageURL: "https://example.com/image2.jpg", Price: 2000},
		{Name: "Test Product 3", ImageURL: "https://example.com/image3.jpg", Price: 3000},
	}
	err := api.PDB.BulkInsert(context.Background(), testProducts)
	if err != nil {
		t.Fatal(err)
	}
}

// TestProductsGet - tests retrieving a product from the database.
func TestProductsGet(t *testing.T) {
	testProductParams := &models.ProductRequestParams{
		Name:     "Test Product",
		ImageURL: "https://example.com/image.jpg",
		Price:    1000,
	}
	addedProduct, err := api.Insert(context.Background(), testProductParams)
	if err != nil {
		t.Fatal(err)
	}
	retrievedProduct, err := api.Get(context.Background(), addedProduct.ID)
	if err != nil {
		t.Fatal(err)
	}
	if *retrievedProduct != *addedProduct {
		t.Errorf("got %v, want %v", *retrievedProduct, *addedProduct)
	}
}

// TestProductsUpdate - tests updating a product in the database.
func TestProductsUpdate(t *testing.T) {
	testProductParams := &models.ProductRequestParams{
		Name:     "Test Product",
		ImageURL: "https://example.com/image.jpg",
		Price:    1000,
	}
	addedProduct, err := api.Insert(context.Background(), testProductParams)
	if err != nil {
		t.Fatal(err)
	}
	ProductRequestParams := &models.ProductRequestParams{
		Name: "Updated Product",
		ImageURL: "https://example.com/updated.jpg",
		Price: 2000,
	}
	updatedProduct, err := api.Update(context.Background(), addedProduct.ID, ProductRequestParams)
	if err != nil {
		t.Fatal(err)
	}
	if updatedProduct.Name != ProductRequestParams.Name || updatedProduct.ImageURL != ProductRequestParams.ImageURL || updatedProduct.Price != ProductRequestParams.Price {
		t.Errorf("got %v, want %v", updatedProduct, ProductRequestParams)
	}
}

// TestProductsDelete - tests deleting a product from the database.
func TestProductsDelete(t *testing.T) {
	testProductParams := &models.ProductRequestParams{
		Name:     "Test Product",
		ImageURL: "https://example.com/image.jpg",
		Price:    1000,
	}
	addedProduct, err := api.Insert(context.Background(), testProductParams)
	if err != nil {
		t.Fatal(err)
	}
	err = api.Delete(context.Background(), addedProduct.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = api.Get(context.Background(), addedProduct.ID)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

// TestProductsGetAll - tests retrieving all products from the database.
func TestProductsGetAll(t *testing.T) {
	testProducts := []*models.ProductRequestParams{
		{Name: "Test Product 1", ImageURL: "https://example.com/image1.jpg", Price: 1000},
		{Name: "Test Product 2", ImageURL: "https://example.com/image2.jpg", Price: 2000},
		{Name: "Test Product 3", ImageURL: "https://example.com/image3.jpg", Price: 3000},
	}
	err := api.PDB.BulkInsert(context.Background(), testProducts)
		if err != nil {
			t.Fatal(err)
		}
	retrievedProducts, err := api.GetAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(retrievedProducts.Data) != len(testProducts) {
		t.Errorf("got %v, want %v", len(retrievedProducts.Data), len(testProducts))
	}
	for i, p := range retrievedProducts.Data {
		if p.Name != testProducts[i].Name || p.ImageURL != testProducts[i].ImageURL || p.Price != testProducts[i].Price {
			t.Errorf("got %v, want %v", p, testProducts[i])
		}
	}
}