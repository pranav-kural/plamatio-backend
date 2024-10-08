CREATE TABLE orders (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id TEXT NOT NULL,
    address_id BIGINT NOT NULL,
    total_price INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    status TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (address_id) REFERENCES addresses(id)
);

CREATE TABLE order_items (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE INDEX idx_user_id_orders ON orders (user_id);

CREATE INDEX idx_order_id_order_items ON order_items (order_id);