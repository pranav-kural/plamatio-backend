CREATE TABLE categories (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    offered BOOLEAN NOT NULL,
);

CREATE TABLE sub_categories (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    category_id BIGINT NOT NULL,
    offered BOOLEAN NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE products (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    category_id BIGINT NOT NULL,
    sub_category_id BIGINT NOT NULL,
    image_url TEXT NOT NULL,
    price INT NOT NULL,
    previous_price INT,
    offered BOOLEAN NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
    FOREIGN KEY (sub_category_id) REFERENCES subcategories(id)
);

CREATE TABLE hero_products (
    product_id BIGINT NOT NULL,
    PRIMARY KEY (product_id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE category_hero_products (
    category_id BIGINT NOT NULL,
    sub_category_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL UNIQUE,
    PRIMARY KEY (category_id, sub_category_id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE users (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
);

CREATE TABLE addresses (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    street TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    country TEXT NOT NULL,
    zip_code TEXT NOT NULL,
    uid TEXT NOT NULL,
    FOREIGN KEY (uid) REFERENCES users(id)
);


CREATE TABLE cart_items (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    uid BIGINT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (uid) REFERENCES users(id)
);


CREATE TABLE orders (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uid BIGINT NOT NULL,
    address_id BIGINT NOT NULL,
    total_price INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    status TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (address_id) REFERENCES address(id)
);

CREATE TABLE order_items (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);