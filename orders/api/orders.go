package orders

import (
	"context"
	"time"

	db "encore.app/orders/db"
	models "encore.app/orders/models"
	rlog "encore.dev/rlog"
	"encore.dev/storage/cache"
	"encore.dev/storage/sqldb"
)

// ------------------------------------------------------
// Setup Database

// Database instance for Plamatio Backend.
var PlamatioDB = sqldb.Named("plamatio_db")

// OrdersTable instance.
var OrdersTable = &db.OrdersTable{DB: PlamatioDB}

// ------------------------------------------------------
// Setup Caching

// OrdersCluster is the cache cluster for orders and order items.
var OrdersCluster = cache.NewCluster("orders-cache-cluster", cache.ClusterConfig{
    // Use LRU policy to evict keys when the cache reaches memory limit.
    EvictionPolicy: cache.AllKeysLRU,
})

// Order Cache Keyspace to store order data by ID.
var OrderCacheKeyspace = cache.NewStructKeyspace[int, models.Order](OrdersCluster, cache.KeyspaceConfig{
	KeyPattern:    "order-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// Orders Cache Keyspace to store all orders for a user.
var UserOrdersCacheKeyspace = cache.NewStructKeyspace[int, models.Orders](OrdersCluster, cache.KeyspaceConfig{
	KeyPattern:    "user-orders-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// ------------------------------------------------------
// Setup API

/*

Primary endpoints for orders:

- GET: /orders/get/:id
- GET: /orders/all/:user_id
- POST: /orders/add
- PUT: /orders/update
- DELETE: /orders/delete/:id

*/

// GET: /orders/get/:id
// Retrieves the order from the database with the given ID.
//encore:api auth method=GET path=/orders/get/:id
func GetOrder(ctx context.Context, id int) (*models.Order, error) {
	// First, try retrieving the order from cache if it exists.
	c, err := OrderCacheKeyspace.Get(ctx, id)
	// if order is found (i.e., no error), return it
	if err == nil {
		return &c, nil
	}
	// If the order is not found in cache, retrieve it from the database.
	r, err := OrdersTable.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to cache the order.
	go func() {
		// Cache the order.
		if err := OrderCacheKeyspace.Set(ctx, id, *r); err != nil {
			// log error
			rlog.Error("Error caching order", err)
		}
	}()
	// Return the order.
	return r, err
}

// GET: /orders/all/:user_id
// Retrieves all orders for a user from the database.
//encore:api auth method=GET path=/orders/all/:user_id
func GetOrders(ctx context.Context, user_id int) (*models.Orders, error) {
	// First, try retrieving all orders for a user from cache if they exist.
	c, err := UserOrdersCacheKeyspace.Get(ctx, user_id)
	// if orders are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the orders are not found in cache, retrieve them from the database.
	r, err := OrdersTable.GetOrdersByUser(ctx, user_id)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to cache the orders.
	go func() {
		// Cache the orders.
		if err := UserOrdersCacheKeyspace.Set(ctx, user_id, *r); err != nil {
			// log error
			rlog.Error("Error caching user orders", err)
		}
	}()
	// Return the orders.
	return r, err
}

// POST: /orders/add
// Inserts an order into the database.
//encore:api auth method=POST path=/orders/add
func AddOrder(ctx context.Context, o *models.OrderRequestParams) (*models.Order, error) {
	// Insert the order into the database.
	or, err := OrdersTable.InsertOrder(ctx, o)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to invalidate the cache for the user's orders.
	go func() {
		// Invalidate the cache for the user's orders.
		_, err = UserOrdersCacheKeyspace.Delete(ctx, o.UserID)
		if err != nil {
			// log error
			rlog.Error("Error deleting user orders cache", err)
		}
	}()

	// TODO: Publish a message to a message broker to notify other services of the change.

	return or, err
}

// PUT: /orders/update
// Updates an order in the database.
//encore:api auth method=PUT path=/orders/update
func UpdateOrder(ctx context.Context, o *models.Order) (*models.OrderRequestStatus, error) {
	// Update the order in the database.
	err := OrdersTable.UpdateOrder(ctx, o)
	if err != nil {
		return &models.OrderRequestStatus{Status: models.OrderRequestFailed}, err
	}
	// Fire a go routine to invalidate the cache for the user's orders.
	go func() {
		// Invalidate the cache for the user's orders.
		_, err = UserOrdersCacheKeyspace.Delete(ctx, o.UserID)
		if err != nil {
			// log error
			rlog.Error("Error deleting user orders cache", err)
		}
	}()

	// TODO: Publish a message to a message broker to notify other services of the change.

	return &models.OrderRequestStatus{Status: models.OrderRequestSuccess}, err
}

// DELETE: /orders/delete/:id
// Deletes an order from the database.
//encore:api auth method=DELETE path=/orders/delete/:id
func DeleteOrder(ctx context.Context, id int) (*models.OrderRequestStatus, error) {
	// Delete the order from the database.
	err := OrdersTable.DeleteOrder(ctx, id)
	if err != nil {
		return &models.OrderRequestStatus{Status: models.OrderRequestFailed}, err
	}
	
	// Fire a go routine to invalidate the cache for the user's orders.
	go func() {
		// Invalidate the cache for the user's orders.
		_, err = UserOrdersCacheKeyspace.Delete(ctx, id)
		if err != nil {
			// log error
			rlog.Error("Error deleting user orders cache", err)
		}
	}()

	// TODO: Publish a message to a message broker to notify other services of the change.
	
	return &models.OrderRequestStatus{Status: models.OrderRequestSuccess}, err
}