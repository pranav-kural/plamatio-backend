package products

import (
	"context"
	"strconv"
	"time"

	db "encore.app/products/db"
	models "encore.app/products/models"
	rlog "encore.dev/rlog"
	"encore.dev/storage/cache"
	"encore.dev/storage/sqldb"
)

// ------------------------------------------------------
// Setup Database

// Database instance for Plamatio Backend.
var PlamatioDB = sqldb.Named("plamatio_db")

// ProductsTB is the products table instance.
var ProductsTB = &db.ProductsTB{DB: PlamatioDB}

// ------------------------------------------------------
// Setup Caching

// ProductsCluster is the cache cluster for products.
var ProductsCluster = cache.NewCluster("products-cache-cluster", cache.ClusterConfig{
    // Use LRU policy to evict keys when the cache reaches memory limit.
    EvictionPolicy: cache.AllKeysLRU,
})

// Product Cache Keyspace to store products by ID.
var ProductCacheKeyspace = cache.NewStructKeyspace[int, models.Product](ProductsCluster, cache.KeyspaceConfig{
	KeyPattern:    "product-cache/:key",
	DefaultExpiry: cache.ExpireIn(24 * time.Hour),
})

// Product Category Cache Keyspace to store products by category.
var ProductCategoryCacheKeyspace = cache.NewStructKeyspace[int, models.Products](ProductsCluster, cache.KeyspaceConfig{
	KeyPattern:    "product-category-cache/:key",
	DefaultExpiry: cache.ExpireIn(24 * time.Hour),
})

// Product Sub-category Cache Keyspace to store products by category.
var ProductSubCategoryCacheKeyspace = cache.NewStructKeyspace[int, models.Products](ProductsCluster, cache.KeyspaceConfig{
	KeyPattern:    "product-sub-category-cache/:key",
	DefaultExpiry: cache.ExpireIn(24 * time.Hour),
})


// Product Search Cache Keyspace to store product search results by query.
var ProductSearchCacheKeyspace = cache.NewStructKeyspace[string, models.Products](ProductsCluster, cache.KeyspaceConfig{
	KeyPattern:    "product-search-cache/:key",
	DefaultExpiry: cache.ExpireIn(24 * time.Hour),
})

// Products Cache Keyspace to store all products at key "all".
var ProductsCacheKeyspace = cache.NewStructKeyspace[string, models.Products](ProductsCluster, cache.KeyspaceConfig{
	KeyPattern:    "products-cache/:key",
	DefaultExpiry: cache.ExpireIn(24 * time.Hour),
})

// Hero Products Cache Keyspace to store hero products.
var HeroProductsCacheKeyspace = cache.NewStructKeyspace[string, models.Products](ProductsCluster, cache.KeyspaceConfig{
	KeyPattern:    "hero-products-cache/:key",
	DefaultExpiry: cache.ExpireIn(24 * time.Hour),
})

// ------------------------------------------------------
// Setup API

// GET: /products/get/:id
// Retrieves the product from the database with the given ID.
//encore:api auth method=GET path=/products/get/:id
func Get(ctx context.Context, id int) (*models.Product, error) {
	// First, try retrieving the product from cache if it exists.
	c, err := ProductCacheKeyspace.Get(ctx, id)
	// if product is found (i.e., no error), return it
	if err == nil {
		return &c, nil
	}
	// If the category is not found in cache, retrieve it from the database.
	// Retrieve the product from the database.
	r, err := ProductsTB.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// Cache the product.
	if err := ProductCacheKeyspace.Set(ctx, id, *r); err != nil {
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
		return &models.Product{
			ID: id, 
			Name: p.Name, 
			Description: p.Description, 
			CategoryId: p.CategoryId, 
			SubCategoryId: p.SubCategoryId, 
			ImageURL: p.ImageURL, 
			Price: p.Price, 
			PreviousPrice: p.PreviousPrice, 
			Offered: p.Offered}, nil
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
	return &models.Product{
			ID: id, 
			Name: p.Name, 
			Description: p.Description, 
			CategoryId: p.CategoryId, 
			SubCategoryId: p.SubCategoryId, 
			ImageURL: p.ImageURL, 
			Price: p.Price, 
			PreviousPrice: p.PreviousPrice, 
			Offered: p.Offered}, nil
}

// GET: /products/all
// Retrieves all products from the database.
//encore:api auth method=GET path=/products/all
func GetAll(ctx context.Context) (*models.Products, error) {
	// First, try retrieving all products from cache if they exist.
	c, err := ProductsCacheKeyspace.Get(ctx, "all")
	// if products are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the products are not found in cache, retrieve them from the database.
	r, err := ProductsTB.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to cache the products.
	go func() {
		// Cache the products.
		if err := ProductsCacheKeyspace.Set(ctx, "all", *r); err != nil {
			// Log the error
			rlog.Error("error caching products data", err)
		}
	}()
	// Return the products.
	return r, err
}

// GET: /products/category/:id
// Retrieves all products from the database by category.
//encore:api auth method=GET path=/products/category/:id
func GetByCategory(ctx context.Context, id int) (*models.Products, error) {
	// First, try retrieving all products from cache if they exist.
	c, err := ProductCategoryCacheKeyspace.Get(ctx, id)
	// if products are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the products are not found in cache, retrieve them from the database.
	r, err := ProductsTB.GetByCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to cache the products.
	go func() {
		// Cache the products.
		if err := ProductCategoryCacheKeyspace.Set(ctx, id, *r); err != nil {
			// Log the error
			rlog.Error("error caching product category data", err)
		}
	}()
	// Return the products.
	return r, err
}

// GET: /products/subcategory/:id
// Retrieves all products from the database by sub-category.
//encore:api auth method=GET path=/products/subcategory/:id
func GetBySubCategory(ctx context.Context, id int) (*models.Products, error) {
	// First, try retrieving all products from cache if they exist.
	c, err := ProductSubCategoryCacheKeyspace.Get(ctx, id)
	// if products are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the products are not found in cache, retrieve them from the database.
	r, err := ProductsTB.GetBySubCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to cache the products.
	go func() {
		// Cache the products.
		if err := ProductSubCategoryCacheKeyspace.Set(ctx, id, *r); err != nil {
			// Log the error
			rlog.Error("error caching product sub-category data", err)
		}
	}()
	// Return the products.
	return r, err
}

// GET: /products/hero
// Retrieves all hero products from the database.
//encore:api auth method=GET path=/products/hero
func GetHeroProducts(ctx context.Context) (*models.Products, error) {
	// First, try retrieving all hero products from cache if they exist.
	c, err := HeroProductsCacheKeyspace.Get(ctx, "all")
	// if hero products are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the hero products are not found in cache, retrieve them from the database.
	r, err := ProductsTB.GetHeroProducts(ctx)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to cache the hero products.
	go func() {
		// Cache the hero products.
		if err := HeroProductsCacheKeyspace.Set(ctx, "all", *r); err != nil {
			// Log the error
			rlog.Error("error caching hero products data", err)
		}
	}()
	// Return the products.
	return r, err
}

// GET: /products/hero/category/:category
// Retrieves all hero products from the database by category.
//encore:api auth method=GET path=/products/hero/category/:category
func GetHeroProductsByCategory(ctx context.Context, category int) (*models.Products, error) {
	c, err := HeroProductsCacheKeyspace.Get(ctx, strconv.Itoa(category))
	// if hero products are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the hero products are not found in cache, retrieve them from the database.
	r, err := ProductsTB.GetCategoryHeroProducts(ctx, category)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to cache the hero products.
	go func() {
		// Cache the hero products by category.
		if err := HeroProductsCacheKeyspace.Set(ctx, strconv.Itoa(category), *r); err != nil {
			// Log the error
			rlog.Error("error caching hero products data", err)
		}
	}()
	// Return the products.
	return r, err
}

// GET: /products/search/:query
// Retrieves all products from the database by search query.
//encore:api auth method=GET path=/products/search/:query
func Search(ctx context.Context, query string) (*models.Products, error) {
	// First, try retrieving all products from cache if they exist.
	c, err := ProductSearchCacheKeyspace.Get(ctx, query)
	// if products are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the products are not found in cache, retrieve them from the database.
	r, err := ProductsTB.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to cache the products.
	go func() {
		// Cache the products.
		if err := ProductSearchCacheKeyspace.Set(ctx, query, *r); err != nil {
			// Log the error
			rlog.Error("error caching product search data", err)
		}
	}()
	// Return the products.
	return r, err
}