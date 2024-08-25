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
	testProduct := &models.Product{
		Name:     "Test Product",
		ImageURL: "https://example.com/image.jpg",
		Price:    1000,
	}
	addedProduct, err := api.Insert(context.Background(), testProduct.Name, testProduct.ImageURL, testProduct.Price)
	if err != nil {
		t.Fatal(err)
	}
	if addedProduct.Name != testProduct.Name || addedProduct.ImageURL != testProduct.ImageURL || addedProduct.Price != testProduct.Price {
		t.Errorf("got %v, want %v", addedProduct, testProduct)
	}
}

// TestProductsGet - tests retrieving a product from the database.
func TestProductsGet(t *testing.T) {
	testProduct := &models.Product{
		Name:     "Test Product",
		ImageURL: "https://example.com/image.jpg",
		Price:    1000,
	}
	addedProduct, err := api.Insert(context.Background(), testProduct.Name, testProduct.ImageURL, testProduct.Price)
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
	testProduct := &models.Product{
		Name:     "Test Product",
		ImageURL: "https://example.com/image.jpg",
		Price:    1000,
	}
	addedProduct, err := api.Insert(context.Background(), testProduct.Name, testProduct.ImageURL, testProduct.Price)
	if err != nil {
		t.Fatal(err)
	}
	newName := "Updated Product"
	newImageURL := "https://example.com/updated.jpg"
	newPrice := 2000
	updatedProduct, err := api.Update(context.Background(), addedProduct.ID, newName, newImageURL, newPrice)
	if err != nil {
		t.Fatal(err)
	}
	if updatedProduct.Name != newName || updatedProduct.ImageURL != newImageURL || updatedProduct.Price != newPrice {
		t.Errorf("got %v, want %v", updatedProduct, &models.Product{Name: newName, ImageURL: newImageURL, Price: newPrice})
	}
}

// TestProductsDelete - tests deleting a product from the database.
func TestProductsDelete(t *testing.T) {
	testProduct := &models.Product{
		Name:     "Test Product",
		ImageURL: "https://example.com/image.jpg",
		Price:    1000,
	}
	addedProduct, err := api.Insert(context.Background(), testProduct.Name, testProduct.ImageURL, testProduct.Price)
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
	testProducts := []*models.Product{
		{Name: "Test Product 1", ImageURL: "https://example.com/image1.jpg", Price: 1000},
		{Name: "Test Product 2", ImageURL: "https://example.com/image2.jpg", Price: 2000},
		{Name: "Test Product 3", ImageURL: "https://example.com/image3.jpg", Price: 3000},
	}
	for _, p := range testProducts {
		_, err := api.Insert(context.Background(), p.Name, p.ImageURL, p.Price)
		if err != nil {
			t.Fatal(err)
		}
	}
	retrievedProducts, err := api.GetAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(retrievedProducts) != len(testProducts) {
		t.Errorf("got %d products, want %d", len(retrievedProducts), len(testProducts))
	}
	for i, p := range retrievedProducts {
		if *p != *testProducts[i] {
			t.Errorf("got %v, want %v", *p, *testProducts[i])
		}
	}
}