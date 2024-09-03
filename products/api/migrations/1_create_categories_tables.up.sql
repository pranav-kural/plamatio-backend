CREATE TABLE categories (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    offered BOOLEAN NOT NULL
);

CREATE TABLE sub_categories (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    category_id BIGINT NOT NULL,
    offered BOOLEAN NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);