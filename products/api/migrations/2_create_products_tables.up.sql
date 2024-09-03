CREATE TABLE products (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    category_id BIGINT NOT NULL,
    sub_category_id BIGINT NOT NULL,
    image_url TEXT NOT NULL,
    price INT NOT NULL,
    previous_price INT,
    offered BOOLEAN NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (sub_category_id) REFERENCES sub_categories(id)
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
    FOREIGN KEY (sub_category_id) REFERENCES sub_categories(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);