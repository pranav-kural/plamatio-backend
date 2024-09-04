package orders

import (
	"context"

	models "encore.app/orders/models"
	rlog "encore.dev/rlog"
)

// ------------------------------------------------------
// Setup API

/*
Primary endpoints to get orders with order items:
- GET: /orders/detailed/get/:id
- GET: /orders/detailed/all/:user_id
*/

// GET: /order/detailed/get/:order_id
// Retrieves the order with order items from the database with the given ID.
//encore:api auth method=GET path=/order/detailed/get/:order_id
func GetDetailedOrder(ctx context.Context, order_id int) (*models.DetailedOrder, error) {
	// Retrieve the order from the database.
	order, err := GetOrder(ctx, order_id)
	if order == nil && err != nil {
		return nil, err
	} else {
		// if we have an error but we also have orders data
		// log error
		rlog.Error("error retrieving order data for detailed order request, but received order data. Likely issue with cache.", err)
	}
	// Retrieve the order items for the order from the database.
	orderItems, err := GetOrderItems(ctx, order_id)
	if orderItems == nil && err != nil {
		return nil, err
	} else {
		// if we have an error but we also have order items data
		// log error
		rlog.Error("error retrieving order items data for detailed order request, but received order items data. Likely issue with cache.", err)
	}
	// Return the detailed order.
	return &models.DetailedOrder{Order: order, Items: orderItems.Data}, nil
}

// GET: /order/detailed/all/:user_id
// Retrieves all orders with order items for a user from the database.
//encore:api auth method=GET path=/order/detailed/all/:user_id
func GetDetailedOrders(ctx context.Context, user_id int) (*models.DetailedOrders, error) {
	// Retrieve all orders for a user from the database.
	orders, err := GetOrders(ctx, user_id)
	if orders == nil && err != nil {
		return nil, err
	} else {
		// if we have an error but we also have orders data
		// log error
		rlog.Error("error retrieving orders data for detailed orders request, but received orders data. Likely issue with cache.", err)
	}
	// Create a new DetailedOrders struct.
	detailedOrders := &models.DetailedOrders{}
	// Loop through each order and retrieve the order items for each order.
	for _, order := range orders.Data {
		orderItems, err := GetOrderItems(ctx, order.ID)
		if orderItems == nil && err != nil {
			return nil, err
		} else {
			// if we have an error but we also have order items data
			// log error
			rlog.Error("error retrieving order items data for detailed orders request, but received order items data. Likely issue with cache.", err)
		}
		// Append the detailed order to the DetailedOrders struct.
		detailedOrders.Data = append(detailedOrders.Data, &models.DetailedOrder{Order: order, Items: orderItems.Data})
	}
	// Return the detailed orders.
	return detailedOrders, nil
}