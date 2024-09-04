package orders

type Order struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	AddressID int    `json:"address_id"`
	TotalPrice int  `json:"total_price"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

type OrderItem struct {
	ID       int `json:"id"`
	OrderID  int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity"`
}

type Orders struct {
	Data []*Order `json:"data"`
}

type OrderItems struct {
	Data []*OrderItem `json:"data"`
}

type DetailedOrder struct {
	Order *Order `json:"order"`
	Items []*OrderItem `json:"items"`
}

type DetailedOrders struct {
	Data []*DetailedOrder `json:"data"`
}

type OrderRequestParams struct {
	UserID    int    `json:"user_id"`
	AddressID int    `json:"address_id"`
	TotalPrice int  `json:"total_price"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

type OrderItemRequestParams struct {
	OrderID  int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity"`
}

type OrderRequestStatus struct {
	Status string `json:"status"`
}

const (
	OrderRequestSuccess = "success"
	OrderRequestFailed = "failed"
)