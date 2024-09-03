-- Insert overall hero products
INSERT INTO hero_products (product_id) VALUES (3);
INSERT INTO hero_products (product_id) VALUES (4);
INSERT INTO hero_products (product_id) VALUES (14);
INSERT INTO hero_products (product_id) VALUES (23);

-- Insert hero products for each category and subcategory
INSERT INTO category_hero_products (category_id, sub_category_id, product_id) VALUES (1, 1, 16);
INSERT INTO category_hero_products (category_id, sub_category_id, product_id) VALUES (1, 2, 8);
INSERT INTO category_hero_products (category_id, sub_category_id, product_id) VALUES (1, 3, 13);
INSERT INTO category_hero_products (category_id, sub_category_id, product_id) VALUES (1, 4, 11);
INSERT INTO category_hero_products (category_id, sub_category_id, product_id) VALUES (2, 5, 1);
INSERT INTO category_hero_products (category_id, sub_category_id, product_id) VALUES (2, 6, 4);
INSERT INTO category_hero_products (category_id, sub_category_id, product_id) VALUES (2, 7, 6);
INSERT INTO category_hero_products (category_id, sub_category_id, product_id) VALUES (3, 8, 21);
INSERT INTO category_hero_products (category_id, sub_category_id, product_id) VALUES (3, 9, 23);
INSERT INTO category_hero_products (category_id, sub_category_id, product_id) VALUES (3, 10, 19);