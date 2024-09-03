package categories

import (
	"context"
	"time"

	models "encore.app/categories/models"
	"encore.dev/storage/cache"
)

// ------------------------------------------------------
// Setup Caching

// Sub-category Cache Keyspace to store category data by ID.
var SubCategoryCacheKeyspace = cache.NewStructKeyspace[int, models.SubCategory](CategoriesCluster, cache.KeyspaceConfig{
	KeyPattern:    "sub-category-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// Sub-categories Cache Keyspace to store all categories at key "all".
var SubCategoriesCacheKeyspace = cache.NewStructKeyspace[string, models.SubCategories](CategoriesCluster, cache.KeyspaceConfig{
	KeyPattern:    "sub-categories-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// ------------------------------------------------------
// Setup API

// GET: /subcategories/get/:id
// Retrieves the sub-category from the database with the given ID.
//encore:api auth method=GET path=/subcategories/get/:id
func GetSubCategory(ctx context.Context, id int) (*models.SubCategory, error) {
	// First, try retrieving the sub-category from cache if it exists.
	c, err := SubCategoryCacheKeyspace.Get(ctx, id)
	// if sub-category is found (i.e., no error), return it
	if err == nil {
		return &c, nil
	}
	// If the sub-category is not found in cache, retrieve it from the database.
	r, err := CategoriesTable.GetSubCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	// Cache the sub-category.
	if err := SubCategoryCacheKeyspace.Set(ctx, id, *r); err != nil {
		return nil, err
	}
	// Return the sub-category.
	return r, err
}

// GET: /subcategories/all
// Retrieves all sub-categories from the database.
//encore:api auth method=GET path=/subcategories/all
func GetSubCategories(ctx context.Context) (*models.SubCategories, error) {
	// First, try retrieving all sub-categories from cache if they exist.
	c, err := SubCategoriesCacheKeyspace.Get(ctx, "all")
	// if sub-categories are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the sub-categories are not found in cache, retrieve them from the database.
	r, err := CategoriesTable.GetAllSubCategories(ctx)
	if err != nil {
		return nil, err
	}
	// Cache the sub-categories.
	if err := SubCategoriesCacheKeyspace.Set(ctx, "all", *r); err != nil {
		return nil, err
	}
	// Return the sub-categories.
	return r, err
}