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

CREATE INDEX idx_category_id ON categories (id);
CREATE INDEX idx_sub_category_id ON sub_categories (id);
CREATE INDEX idx_category_id_sub_categories ON sub_categories (category_id);