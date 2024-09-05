package categories

import (
	"context"
	"time"

	db "encore.app/categories/db"
	models "encore.app/categories/models"
	rlog "encore.dev/rlog"
	"encore.dev/storage/cache"
	"encore.dev/storage/sqldb"
)

// ------------------------------------------------------
// Setup Database

// ProductDB instance.
var ProductsDB = sqldb.Named("products")

// CategoriesTable instance.
var CategoriesTable = &db.CategoriesTable{DB: ProductsDB}

// ------------------------------------------------------
// Setup Caching

// CategoriesCluster is the cache cluster for categories.
var CategoriesCluster = cache.NewCluster("categories-cache-cluster", cache.ClusterConfig{
    // Use LRU policy to evict keys when the cache reaches memory limit.
    EvictionPolicy: cache.AllKeysLRU,
})

// Category Cache Keyspace to store category data by ID.
var CategoryCacheKeyspace = cache.NewStructKeyspace[int, models.Category](CategoriesCluster, cache.KeyspaceConfig{
	KeyPattern:    "category-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// Categories Cache Keyspace to store all categories at key "all".
var CategoriesCacheKeyspace = cache.NewStructKeyspace[string, models.Categories](CategoriesCluster, cache.KeyspaceConfig{
	KeyPattern:    "categories-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// ------------------------------------------------------
// Setup API

// GET: /categories/get/:id
// Retrieves the category from the database with the given ID.
//encore:api auth method=GET path=/categories/get/:id
func GetCategory(ctx context.Context, id int) (*models.Category, error) {
	// First, try retrieving the category from cache if it exists.
	c, err := CategoryCacheKeyspace.Get(ctx, id)
	// if category is found (i.e., no error), return it
	if err == nil {
		return &c, nil
	}
	// If the category is not found in cache, retrieve it from the database.
	r, err := CategoriesTable.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// Fire go routine to cache the category.
	go func() {
		// Cache the category.
		if err := CategoryCacheKeyspace.Set(ctx, id, *r); err != nil {
			// log error
			rlog.Error("Error caching category", err)
		}
	}()
	// Return the category.
	return r, err
}

// GET: /categories/all
// Retrieves all categories from the database.
//encore:api auth method=GET path=/categories/all
func GetCategories(ctx context.Context) (*models.Categories, error) {
	// First, try retrieving all categories from cache if they exist.
	c, err := CategoriesCacheKeyspace.Get(ctx, "all")
	// if categories are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the categories are not found in cache, retrieve them from the database.
	r, err := CategoriesTable.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	// Fire go routine to cache the categories.
	go func() {
		// Cache the categories.
		if err := CategoriesCacheKeyspace.Set(ctx, "all", *r); err != nil {
			// log error
			rlog.Error("Error caching categories", err)
		}
	}()
	// Return the categories.
	return r, err
}