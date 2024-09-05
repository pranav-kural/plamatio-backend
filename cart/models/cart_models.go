package cart

// CartItem represents an item in the cart.
type CartItem struct {
	ID        int `json:"id"`          // ID is the unique identifier of the cart item.
	ProductID int `json:"product_id"`  // ProductID is the identifier of the product associated with the cart item.
	Quantity  int   `json:"quantity"`   // Quantity is the number of items in the cart.
	UserID    int `json:"user_id"`     // UserID is the identifier of the user who owns the cart item.
}

// CartItems represents a collection of cart items.
type CartItems struct {
	Data []*CartItem `json:"data"`  // Data is the list of cart items.
}

// NewCartItem represents a new cart item to be added to the cart.
type NewCartItem struct {
	ProductID int `json:"product_id"`  // ProductID is the identifier of the product to be added to the cart.
	Quantity  int `json:"quantity"`    // Quantity is the number of items to be added to the cart.
	UserID    int `json:"user_id"`     // UserID is the identifier of the user who owns the cart.
}

// NewCartItems represents a collection of new cart items to be added to the cart.
type NewCartItems struct {
	Data []*NewCartItem `json:"data"`  // Data is the list of new cart items.
}

// CartChangeRequestStatus represents the status of a cart change request.
type CartChangeRequestStatus struct {
	Status string `json:"status"`  // Status is the status of the cart change request.
}

// CartRequestSuccess is the status for a successful cart request.
const CartRequestSuccess = "success"

// CartRequestFailed is the status for a failed cart request.
const CartRequestFailed = "failed"