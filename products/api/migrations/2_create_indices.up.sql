CREATE INDEX idx_category_id ON categories (id);

CREATE INDEX idx_sub_category_id ON sub_categories (id);
CREATE INDEX idx_category_id_sub_categories ON sub_categories (category_id);

CREATE INDEX idx_product_id ON products (id);
CREATE INDEX idx_category_id_products ON products (category_id);
CREATE INDEX idx_sub_category_id_products ON products (sub_category_id);

CREATE INDEX idx_category_id_category_hero_products ON category_hero_products (category_id);

CREATE INDEX idx_uid_users ON users (id);

CREATE INDEX idx_uid_addresses ON addresses (uid);

CREATE INDEX idx_uid_cart_items ON cart_items (uid);

CREATE INDEX idx_uid_orders ON orders (uid);

CREATE INDEX idx_order_id_order_items ON order_items (order_id);