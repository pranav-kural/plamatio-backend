package cart

import (
	"context"

	models "encore.app/cart/models"
	"encore.dev/storage/sqldb"
)

type CartItemsTable struct {
	DB *sqldb.Database
}

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

CREATE INDEX idx_user_id_cart_items ON cart_items (user_id);

*/

const (
    SQL_GET_CART_ITEM = `
				SELECT product_id, quantity, user_id FROM cart_items
				WHERE id = $1
		`
		SQL_GET_ALL_CART_ITEMS = `
				SELECT id, product_id, quantity, user_id FROM cart_items
		`
		SQL_GET_CART_ITEMS_BY_USER = `
				SELECT id, product_id, quantity, user_id FROM cart_items
				WHERE user_id = $1
		`
		SQL_INSERT_CART_ITEM = `
				INSERT INTO cart_items (product_id, quantity, user_id) VALUES ($1, $2, $3) RETURNING id
		`
		SQL_UPDATE_CART_ITEM = `
				UPDATE cart_items SET product_id = $1, quantity = $2, user_id = $3 WHERE id = $4
		`
		SQL_DELETE_CART_ITEM = `
				DELETE FROM cart_items WHERE id = $1
		`
)

// Retrieves a cart item from the database.
func (tb *CartItemsTable) GetCartItem(ctx context.Context, id int) (*models.CartItem, error) {
	ci := &models.CartItem{ID: id}
	err := tb.DB.QueryRow(ctx, SQL_GET_CART_ITEM, id).Scan(&ci.ProductID, &ci.Quantity, &ci.UserID)
	return ci, err
}

// Retrieves all cart items from the database.
func (tb *CartItemsTable) GetAllCartItems(ctx context.Context) (*models.CartItems, error) {
	rows, err := tb.DB.Query(ctx, SQL_GET_ALL_CART_ITEMS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []*models.CartItem
	for rows.Next() {
		ci := &models.CartItem{}
		if err := rows.Scan(&ci.ID, &ci.ProductID, &ci.Quantity, &ci.UserID); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, ci)
	}
	return &models.CartItems{Data: cartItems}, nil
}

// Retrieves all cart items for a user from the database.
func (tb *CartItemsTable) GetCartItemsByUser(ctx context.Context, userId int) (*models.CartItems, error) {
	rows, err := tb.DB.Query(ctx, SQL_GET_CART_ITEMS_BY_USER, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []*models.CartItem
	for rows.Next() {
		ci := &models.CartItem{}
		if err := rows.Scan(&ci.ID, &ci.ProductID, &ci.Quantity, &ci.UserID); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, ci)
	}
	return &models.CartItems{Data: cartItems}, nil
}

// Inserts a cart item into the database.
func (tb *CartItemsTable) InsertCartItem(ctx context.Context, productID int, quantity int, userID int) (int, error) {
	var id int
	err := tb.DB.QueryRow(ctx, SQL_INSERT_CART_ITEM, productID, quantity, userID).Scan(&id)
	return id, err
}

// Updates a cart item in the database.
func (tb *CartItemsTable) UpdateCartItem(ctx context.Context, productID int, quantity int, userID int, id int) error {
	_, err := tb.DB.Exec(ctx, SQL_UPDATE_CART_ITEM, productID, quantity, userID, id)
	return err
}

// Deletes a cart item from the database.
func (tb *CartItemsTable) DeleteCartItem(ctx context.Context, id int) error {
	_, err := tb.DB.Exec(ctx, SQL_DELETE_CART_ITEM, id)
	return err
}