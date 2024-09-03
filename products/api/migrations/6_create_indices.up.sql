CREATE INDEX idx_category_id ON categories (id);

CREATE INDEX idx_sub_category_id ON sub_categories (id);
CREATE INDEX idx_category_id_sub_categories ON sub_categories (category_id);

CREATE INDEX idx_product_id ON products (id);
CREATE INDEX idx_category_id_products ON products (category_id);
CREATE INDEX idx_sub_category_id_products ON products (sub_category_id);

CREATE INDEX idx_category_id_category_hero_products ON category_hero_products (category_id);

CREATE INDEX idx_user_id_users ON users (id);

CREATE INDEX idx_user_id_addresses ON addresses (user_id);

CREATE INDEX idx_user_id_cart_items ON cart_items (user_id);

CREATE INDEX idx_user_id_orders ON orders (user_id);

CREATE INDEX idx_order_id_order_items ON order_items (order_id);