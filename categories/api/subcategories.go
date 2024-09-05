package categories

import (
	"context"
	"time"

	db "encore.app/categories/db"
	models "encore.app/categories/models"
	rlog "encore.dev/rlog"
	"encore.dev/storage/cache"
)

// ------------------------------------------------------
// Setup Database
// SubCategoriesTable instance.
var SubCategoriesTable = &db.SubCategoriesTable{DB: ProductsDB}

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

// Category Sub-categories Cache Keyspace to store sub-categories by category ID.
var SubCategoriesByCategoryCacheKeyspace = cache.NewStructKeyspace[int, models.SubCategories](CategoriesCluster, cache.KeyspaceConfig{
	KeyPattern:    "sub-categories-by-category-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// ------------------------------------------------------
// Setup API

// GET: /categories/subcategories/get/:id
// Retrieves the sub-category from the database with the given ID.
//encore:api auth method=GET path=/categories/subcategories/get/:id
func GetSubCategory(ctx context.Context, id int) (*models.SubCategory, error) {
	// First, try retrieving the sub-category from cache if it exists.
	c, err := SubCategoryCacheKeyspace.Get(ctx, id)
	// if sub-category is found (i.e., no error), return it
	if err == nil {
		return &c, nil
	}
	// If the sub-category is not found in cache, retrieve it from the database.
	r, err := SubCategoriesTable.GetSubCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	// Cache the sub-category.
	if err := SubCategoryCacheKeyspace.Set(ctx, id, *r); err != nil {
		// log error
		rlog.Error("Error caching sub-category", err)
	}
	// Return the sub-category.
	return r, err
}

// GET: /categories/subcategories/all
// Retrieves all sub-categories from the database.
//encore:api auth method=GET path=/categories/subcategories/all
func GetSubCategories(ctx context.Context) (*models.SubCategories, error) {
	// First, try retrieving all sub-categories from cache if they exist.
	c, err := SubCategoriesCacheKeyspace.Get(ctx, "all")
	// if sub-categories are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the sub-categories are not found in cache, retrieve them from the database.
	r, err := SubCategoriesTable.GetAllSubCategories(ctx)
	if err != nil {
		return nil, err
	}
	// Cache the sub-categories.
	if err := SubCategoriesCacheKeyspace.Set(ctx, "all", *r); err != nil {
		// log error
		rlog.Error("Error caching sub-categories", err)
	}
	// Return the sub-categories.
	return r, err
}

// GET: /categories/subcategories/category/:id
// Retrieves all sub-categories from the database by category.
//encore:api auth method=GET path=/categories/subcategories/category/:id
func GetSubCategoriesByCategory(ctx context.Context, categoryId int) (*models.SubCategories, error) {
	// First, try retrieving all sub-categories from cache if they exist.
	c, err := SubCategoriesByCategoryCacheKeyspace.Get(ctx, categoryId)
	// if sub-categories are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the sub-categories are not found in cache, retrieve them from the database.
	r, err := SubCategoriesTable.GetSubCategoriesByCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	// Cache the sub-categories.
	if err := SubCategoriesByCategoryCacheKeyspace.Set(ctx, categoryId, *r); err != nil {
		// log error
		rlog.Error("Error caching sub-categories by category", err)
	}
	// Return the sub-categories.
	return r, err
}