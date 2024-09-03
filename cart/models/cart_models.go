package cart

/*
For reference, here is the SQL to create the tables in the database:

CREATE TABLE cart_items (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    user_id BIGINT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

*/

type CartItem struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity  int   `json:"quantity"`
	UserID    int `json:"user_id"`
}

type CartItems struct {
	Data []*CartItem `json:"data"`
}