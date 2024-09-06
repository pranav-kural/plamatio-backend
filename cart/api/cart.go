package cart

import (
	"context"
	"errors"
	"time"

	db "encore.app/cart/db"
	models "encore.app/cart/models"
	rlog "encore.dev/rlog"
	"encore.dev/storage/cache"
	"encore.dev/storage/sqldb"
)

// ------------------------------------------------------
// Setup Database

// Database instance for Plamatio Backend.
var PlamatioDB = sqldb.Named("plamatio_db")

// CartItemsTable instance.
var CartItemsTable = &db.CartItemsTable{DB: PlamatioDB}

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
	// Fire go routine to cache the cart item.
	go func() {
		// Cache the cart item.
		if err := CartItemCacheKeyspace.Set(ctx, id, *r); err != nil {
			// log error
			rlog.Error("Error caching cart item", err)
		}
	}()
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
	// Fire go routine to cache the cart items.
	go func() {
		// Cache the cart items.
		if err := CartItemsCacheKeyspace.Set(ctx, user_id, *r); err != nil {
			// log error
			rlog.Error("Error caching user cart items", err)
		}
	}()
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
	// Fire go routine to invalidate the cache for the user's cart items.
	go func() {
		// Invalidate the cache for the user's cart items.
		_, err = CartItemsCacheKeyspace.Delete(ctx, newCartItem.UserID)
		if err != nil {
			// log error
			rlog.Error("Error deleting user cart items cache", err)
		}
	}()

	// TODO: Send event on Kafka topic for cart items update for the user.

	return r, err
}

// PUT: /cart/update
// Updates a cart item in the database.
//encore:api auth method=PUT path=/cart/update
func UpdateCartItem(ctx context.Context, updatedCartItem *models.CartItem) (*models.CartChangeRequestReturn, error) {
	// Update the cart item in the database.
	err := CartItemsTable.UpdateCartItem(ctx, updatedCartItem.ProductID, updatedCartItem.Quantity, updatedCartItem.UserID, updatedCartItem.ID)
	if err != nil {
		return nil, err
	}
	// Fire go routine to invalidate the cache for the user's cart items.
	go func() {
		// Invalidate the cache for the user's cart items.
		_, err = CartItemsCacheKeyspace.Delete(ctx, updatedCartItem.UserID)
		if err != nil {
			// log error
			rlog.Error("Error deleting user cart items cache", err)
		}
	}()

	// TODO: Send event on Kafka topic for cart items update for the user.

	return &models.CartChangeRequestReturn{CartID: updatedCartItem.ID}, nil
}

// DELETE: /cart/delete/:id
// Deletes a cart item from the database.
//encore:api auth method=DELETE path=/cart/delete/:id
func DeleteCartItem(ctx context.Context, id int) (*models.CartChangeRequestReturn, error) {
	// Delete the cart item from the database.
	err := CartItemsTable.DeleteCartItem(ctx, id)
	if err != nil {
		return nil, err
	}
	// Fire go routine to invalidate the cache for the user's cart items.
	go func() {
		// Invalidate the cache for the user's cart items.
		_, err = CartItemsCacheKeyspace.Delete(ctx, id)
		if err != nil {
			// log error
			rlog.Error("Error deleting user cart items cache", err)
		}
	}()

	// TODO: Send event on Kafka topic for cart items update for the user.

	return &models.CartChangeRequestReturn{CartID: id}, nil
}