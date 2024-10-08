// Package orders provides models for handling orders and order-related data.
package orders

import "time"

// Order represents an order entity.
type Order struct {
	ID        int    `json:"id"`          // Unique identifier for the order.
	UserID    string    `json:"user_id"`     // ID of the user who placed the order.
	AddressID int    `json:"address_id"`  // ID of the address associated with the order.
	TotalPrice float64  `json:"total_price"`  // Total price of the order.
	CreatedAt time.Time `json:"created_at"`  // Timestamp indicating when the order was created.
	Status    string `json:"status"`      // Current status of the order.
}

// OrderItem represents an item within an order.
type OrderItem struct {
	ID        int `json:"id"`             // Unique identifier for the order item.
	OrderID   int `json:"order_id"`       // ID of the order to which the item belongs.
	ProductID int `json:"product_id"`     // ID of the product associated with the item.
	Quantity  int `json:"quantity"`       // Quantity of the item.
}

// Orders represents a collection of orders.
type Orders struct {
	Data []*Order `json:"data"`           // List of order entities.
}

// OrderItems represents a collection of order items.
type OrderItems struct {
	Data []*OrderItem `json:"data"`       // List of order item entities.
}

// DetailedOrder represents an order with its associated items.
type DetailedOrder struct {
	Order *Order         `json:"order"`    // The order entity.
	Items []*OrderItem   `json:"items"`    // List of order item entities.
}

// DetailedOrders represents a collection of detailed orders.
type DetailedOrders struct {
	Data []*DetailedOrder `json:"data"`     // List of detailed order entities.
}

// OrderRequestParams represents the parameters for creating or updating an order.
type OrderRequestParams struct {
	UserID    string    `json:"user_id"`      // ID of the user placing the order.
	AddressID int    `json:"address_id"`   // ID of the address associated with the order.
	TotalPrice float64  `json:"total_price"`   // Total price of the order.
	Status    string `json:"status"`       // Current status of the order.
}

// DetailedOrderItemRequestParams represents the parameters for creating or updating an order item with product details.
type DetailedOrderItemRequestParams struct {
	ProductID int `json:"product_id"`      // ID of the product associated with the item.
	Quantity  int `json:"quantity"`        // Quantity of the item.
}

// DetailedOrderRequestParams represents the parameters for creating or updating an order with items.
type DetailedOrderRequestParams struct {
	Order *OrderRequestParams `json:"order"` // The order entity.
	Items []*DetailedOrderItemRequestParams `json:"items"` // List of order item entities.
}

// OrderItemRequestParams represents the parameters for creating or updating an order item.
type OrderItemRequestParams struct {
	OrderID   int `json:"order_id"`        // ID of the order to which the item belongs.
	ProductID int `json:"product_id"`      // ID of the product associated with the item.
	Quantity  int `json:"quantity"`        // Quantity of the item.
}

// Order mutation request return type.
type OrderChangeRequestReturn struct {
	OrderID int `json:"id"`                // ID of the order.
}

// OrderItem mutation request return type.
type OrderItemChangeRequestReturn struct {
	OrderItemID int `json:"id"`            // ID of the order item.
}