CREATE TABLE products (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    description TEXT,
    category TEXT NOT NULL,
    imageUrl TEXT NOT NULL,
    price INT NOT NULL
  );