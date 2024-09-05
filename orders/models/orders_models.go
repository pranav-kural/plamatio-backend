// Package orders provides models for handling orders and order-related data.
package orders

// Order represents an order entity.
type Order struct {
	ID        int    `json:"id"`          // Unique identifier for the order.
	UserID    int    `json:"user_id"`     // ID of the user who placed the order.
	AddressID int    `json:"address_id"`  // ID of the address associated with the order.
	TotalPrice int  `json:"total_price"`  // Total price of the order.
	CreatedAt string `json:"created_at"`  // Timestamp indicating when the order was created.
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
	UserID    int    `json:"user_id"`      // ID of the user placing the order.
	AddressID int    `json:"address_id"`   // ID of the address associated with the order.
	TotalPrice int  `json:"total_price"`   // Total price of the order.
	CreatedAt string `json:"created_at"`   // Timestamp indicating when the order was created.
	Status    string `json:"status"`       // Current status of the order.
}

// OrderItemRequestParams represents the parameters for creating or updating an order item.
type OrderItemRequestParams struct {
	OrderID   int `json:"order_id"`        // ID of the order to which the item belongs.
	ProductID int `json:"product_id"`      // ID of the product associated with the item.
	Quantity  int `json:"quantity"`        // Quantity of the item.
}

// OrderRequestStatus represents the status of an order request.
type OrderRequestStatus struct {
	Status string `json:"status"`          // Status of the order request.
}

// Constants for order request status.
const (
	OrderRequestSuccess = "success"         // Order request was successful.
	OrderRequestFailed  = "failed"          // Order request failed.
)