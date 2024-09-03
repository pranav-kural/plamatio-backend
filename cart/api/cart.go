package cart

import (
	"context"
	"errors"
	"time"

	db "encore.app/cart/db"
	models "encore.app/cart/models"
	"encore.dev/storage/cache"
	"encore.dev/storage/sqldb"
)

// ------------------------------------------------------
// Setup Database

// ProductDB instance.
var ProductsDB = sqldb.Named("products")

// CartItemsTable instance.
var CartItemsTable = &db.CartItemsTable{DB: ProductsDB}

// ------------------------------------------------------
// Setup Caching

// CartCluster is the cache cluster for cart items.
var CartCluster = cache.NewCluster("cart-cache-cluster", cache.ClusterConfig{
    // Use LRU policy to evict keys when the cache reaches memory limit.
    EvictionPolicy: cache.AllKeysLRU,
})

// Cart Item Cache Keyspace to store cart items data by ID.
var CartItemCacheKeyspace = cache.NewStructKeyspace[int, models.CartItem](CartCluster, cache.KeyspaceConfig{
	KeyPattern:    "cart-item-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// Cart Items Cache Keyspace to store all cart items for a user.
var CartItemsCacheKeyspace = cache.NewStructKeyspace[int, models.CartItems](CartCluster, cache.KeyspaceConfig{
	KeyPattern:    "user-cart-items-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// ------------------------------------------------------
// Setup API

// GET: /cart/get/:id
// Retrieves the cart item from the database with the given ID.
//encore:api auth method=GET path=/cart/get/:id
func GetCartItem(ctx context.Context, id int) (*models.CartItem, error) {
	// First, try retrieving the cart item from cache if it exists.
	c, err := CartItemCacheKeyspace.Get(ctx, id)
	// if cart item is found (i.e., no error), return it
	if err == nil {
		return &c, nil
	}
	// If the cart item is not found in cache, retrieve it from the database.
	r, err := CartItemsTable.GetCartItem(ctx, id)
	if err != nil {
		return nil, err
	}
	// Cache the cart item.
	if err := CartItemCacheKeyspace.Set(ctx, id, *r); err != nil {
		return nil, err
	}
	// Return the cart item.
	return r, err
}

// GET: /cart/all/:user_id
// Retrieves all cart items for a user from the database.
//encore:api auth method=GET path=/cart/all/:user_id
func GetCartItems(ctx context.Context, user_id int) (*models.CartItems, error) {
	// confirm user_id is valid - less than 1
	if user_id < 1 {
		return nil, errors.New("invalid user_id")
	}
	// First, try retrieving all cart items for a user from cache if they exist.
	c, err := CartItemsCacheKeyspace.Get(ctx, user_id)
	// if cart items are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the cart items are not found in cache, retrieve them from the database.
	r, err := CartItemsTable.GetCartItemsByUser(ctx, user_id)
	if err != nil {
		return nil, err
	}
	// Cache the cart items.
	if err := CartItemsCacheKeyspace.Set(ctx, user_id, *r); err != nil {
		return nil, err
	}
	// Return the cart items.
	return r, err
}

// POST: /cart/add
// Inserts a cart item into the database.
//encore:api auth method=POST path=/cart/add
func AddCartItem(ctx context.Context, newCartItem *models.NewCartItem) (*models.CartItem, error) {
	// Insert the cart item into the database.
	r, err := CartItemsTable.InsertCartItem(ctx, newCartItem.ProductID, newCartItem.Quantity, newCartItem.UserID)
	if err != nil {
		return nil, err
	}
	// Invalidate the cache for the user's cart items.
	_, err = CartItemsCacheKeyspace.Delete(ctx, newCartItem.UserID)
	if err != nil {
		return nil, err
	}

	// TODO: Send event on Kafka topic for cart items update for the user.

	return r, nil
}

// PUT: /cart/update
// Updates a cart item in the database.
//encore:api auth method=PUT path=/cart/update
func UpdateCartItem(ctx context.Context, updatedCartItem *models.CartItem) (*models.CartChangeRequestStatus, error) {
	// Update the cart item in the database.
	err := CartItemsTable.UpdateCartItem(ctx, updatedCartItem.ProductID, updatedCartItem.Quantity, updatedCartItem.UserID, updatedCartItem.ID)
	if err != nil {
		return &models.CartChangeRequestStatus{Status: models.CartRequestFailed}, err
	}
	// Invalidate the cache for the user's cart items.
	_, err = CartItemsCacheKeyspace.Delete(ctx, updatedCartItem.UserID)
	if err != nil {
		return &models.CartChangeRequestStatus{Status: models.CartRequestFailed}, err
	}

	// TODO: Send event on Kafka topic for cart items update for the user.

	return &models.CartChangeRequestStatus{Status: models.CartRequestSuccess}, nil
}

// DELETE: /cart/delete/:id
// Deletes a cart item from the database.
//encore:api auth method=DELETE path=/cart/delete/:id
func DeleteCartItem(ctx context.Context, id int) (*models.CartChangeRequestStatus, error) {
	// Delete the cart item from the database.
	err := CartItemsTable.DeleteCartItem(ctx, id)
	if err != nil {
		return &models.CartChangeRequestStatus{Status: models.CartRequestFailed}, err
	}
	// Invalidate the cache for the user's cart items.
	_, err = CartItemsCacheKeyspace.Delete(ctx, id)
	if err != nil {
		return &models.CartChangeRequestStatus{Status: models.CartRequestFailed}, err
	}

	// TODO: Send event on Kafka topic for cart items update for the user.

	return &models.CartChangeRequestStatus{Status: models.CartRequestSuccess}, nil
}