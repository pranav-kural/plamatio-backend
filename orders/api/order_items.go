package orders

import (
	"context"
	"time"

	db "encore.app/orders/db"
	models "encore.app/orders/models"
	rlog "encore.dev/rlog"
	"encore.dev/storage/cache"
)

// ------------------------------------------------------
// Setup Database

// OrderItemsTable instance.
var OrderItemsTable = &db.OrderItemsTable{DB: PlamatioDB}

// ------------------------------------------------------
// Setup Caching

// Order Item Cache Keyspace to store order items data by ID.
var OrderItemCacheKeyspace = cache.NewStructKeyspace[int, models.OrderItem](OrdersCluster, cache.KeyspaceConfig{
	KeyPattern:    "order-item-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// Order Items Cache Keyspace to store all order items for a specific order by order ID.
var OrderItemsCacheKeyspace = cache.NewStructKeyspace[int, models.OrderItems](OrdersCluster, cache.KeyspaceConfig{
	KeyPattern:    "order-items-cache/:key",
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

// GET: /orders/items/get/:id
// Retrieves the order item from the database with the given ID.
//encore:api auth method=GET path=/orders/items/get/:id
func GetOrderItem(ctx context.Context, id int) (*models.OrderItem, error) {
	// First, try retrieving the order item from cache if it exists.
	c, err := OrderItemCacheKeyspace.Get(ctx, id)
	// if order item is found (i.e., no error), return it
	if err == nil {
		return &c, nil
	}
	// If the order item is not found in cache, retrieve it from the database.
	r, err := OrderItemsTable.GetOrderItem(ctx, id)
	if err != nil {
		return nil, err
	}
	// Fire go routine to cache the order item.
	go func() {
		// Cache the order item.
		if err := OrderItemCacheKeyspace.Set(ctx, id, *r); err != nil {
			// log error
			rlog.Error("Error caching order item", err)
		}
	}()
	// Return the order item.
	return r, err
}

// GET: /orders/items/all/:order_id
// Retrieves all order items for an order from the database.
//encore:api auth method=GET path=/orders/items/all/:order_id
func GetOrderItems(ctx context.Context, order_id int) (*models.OrderItems, error) {
	// First, try retrieving all order items for an order from cache if they exist.
	c, err := OrderItemsCacheKeyspace.Get(ctx, order_id)
	// if order items are found (i.e., no error), return them
	if err == nil {
		return &c, nil
	}
	// If the order items are not found in cache, retrieve them from the database.
	r, err := OrderItemsTable.GetOrderItemsByOrder(ctx, order_id)
	if err != nil {
		return nil, err
	}
	// Fire go routine to cache the order items.
	go func() {
		// Cache the order items.
		if err := OrderItemsCacheKeyspace.Set(ctx, order_id, *r); err != nil {
			// log error
			rlog.Error("Error caching order items", err)
		}
	}()
	// Return the order items.
	return r, err
}

// POST: /orders/items/add
// Inserts an order item into the database.
//encore:api auth method=POST path=/orders/items/add
func AddOrderItem(ctx context.Context, oi *models.OrderItemRequestParams) (*models.OrderItem, error) {
	// Insert the order item into the database.
	noi, err := OrderItemsTable.InsertOrderItem(ctx, oi)
	if err != nil {
		return nil, err
	}
	// Fire go routine to invalidate the cache for the order's order items.
	go func() {
		// Invalidate the cache for the order's order items.
		_, err = OrderItemsCacheKeyspace.Delete(ctx, oi.OrderID)
		if err != nil {
			// log error
			rlog.Error("Error deleting order items cache", err)
		}
	}()

	// TODO: Publish a message to a message broker to notify other services of the change.

	return noi, err
}

// PUT: /orders/items/update
// Updates an order item in the database.
//encore:api auth method=PUT path=/orders/items/update
func UpdateOrderItem(ctx context.Context, oi *models.OrderItem) (*models.OrderRequestStatus, error) {
	// Update the order item in the database.
	err := OrderItemsTable.UpdateOrderItem(ctx, oi)
	if err != nil {
		return &models.OrderRequestStatus{Status: models.OrderRequestFailed}, err
	}
	// Fire go routine to invalidate the cache for the order's order items.
	go func() {
		// Invalidate the cache for the order's order items.
		_, err = OrderItemsCacheKeyspace.Delete(ctx, oi.OrderID)
		if err != nil {
			// log error
			rlog.Error("Error deleting order items cache", err)
		}
	}()

	// TODO: Publish a message to a message broker to notify other services of the change.

	return &models.OrderRequestStatus{Status: models.OrderRequestSuccess}, err
}

// DELETE: /orders/items/delete/:id
// Deletes an order item from the database.
//encore:api auth method=DELETE path=/orders/items/delete/:id
func DeleteOrderItem(ctx context.Context, id int) (*models.OrderRequestStatus, error) {
	// Delete the order item from the database.
	err := OrderItemsTable.DeleteOrderItem(ctx, id)
	if err != nil {
		return &models.OrderRequestStatus{Status: models.OrderRequestFailed}, err
	}
	// Fire go routine to invalidate the cache for the order's order items.
	go func() {
		// Invalidate the cache for the order's order items.
		_, err = OrderItemsCacheKeyspace.Delete(ctx, id)
		if err != nil {
			// log error
			rlog.Error("Error deleting order items cache", err)
		}
	}()

	// TODO: Publish a message to a message broker to notify other services of the change.
	
	return &models.OrderRequestStatus{Status: models.OrderRequestSuccess}, err
}