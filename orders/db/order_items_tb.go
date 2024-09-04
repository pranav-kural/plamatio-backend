package orders

import (
	"context"

	models "encore.app/orders/models"
	utils "encore.app/orders/utils"
	"encore.dev/storage/sqldb"
)

type OrderItemsTable struct {
	DB *sqldb.Database
}

const (
		SQL_GET_ORDER_ITEM = `
				SELECT product_id, quantity FROM order_items
				WHERE id = $1
		`
		SQL_GET_ALL_ORDER_ITEMS = `
				SELECT id, product_id, quantity FROM order_items
		`
		SQL_GET_ORDER_ITEMS_BY_ORDER = `
				SELECT id, product_id, quantity FROM order_items
				WHERE order_id = $1
		`
		SQL_INSERT_ORDER_ITEM = `
				INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING id
		`
		SQL_UPDATE_ORDER_ITEM = `
				UPDATE order_items SET order_id = $1, product_id = $2, quantity = $3 WHERE id = $4
		`
		SQL_DELETE_ORDER_ITEM = `
				DELETE FROM order_items WHERE id = $1
		`
)

// Retrieves an order item from the database.
func (tb *OrderItemsTable) GetOrderItem(ctx context.Context, id int) (*models.OrderItem, error) {
	oi := &models.OrderItem{ID: id}
	err := tb.DB.QueryRow(ctx, SQL_GET_ORDER_ITEM, id).Scan(&oi.ProductID, &oi.Quantity)
	return oi, err
}

// Retrieves all order items from the database.
func (tb *OrderItemsTable) GetAllOrderItems(ctx context.Context) (*models.OrderItems, error) {
	rows, err := tb.DB.Query(ctx, SQL_GET_ALL_ORDER_ITEMS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderItems := &models.OrderItems{}
	for rows.Next() {
		oi := &models.OrderItem{}
		if err := rows.Scan(&oi.ID, &oi.ProductID, &oi.Quantity); err != nil {
			return nil, err
		}
		orderItems.Data = append(orderItems.Data, oi)
	}
	return orderItems, nil
}

// Retrieves all order items for an order from the database.
func (tb *OrderItemsTable) GetOrderItemsByOrder(ctx context.Context, orderId int) (*models.OrderItems, error) {
	rows, err := tb.DB.Query(ctx, SQL_GET_ORDER_ITEMS_BY_ORDER, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderItems := &models.OrderItems{}
	for rows.Next() {
		oi := &models.OrderItem{}
		if err := rows.Scan(&oi.ID, &oi.ProductID, &oi.Quantity); err != nil {
			return nil, err
		}
		orderItems.Data = append(orderItems.Data, oi)
	}
	return orderItems, nil
}

// Inserts an order item into the database.
func (tb *OrderItemsTable) InsertOrderItem(ctx context.Context, oi *models.OrderItemRequestParams) (*models.OrderItem, error) {
	// validate data
	if err := utils.ValidateNewOrderItemData(oi); err != nil {
		return nil, err
	}
	var oiID int
	err := tb.DB.QueryRow(ctx, SQL_INSERT_ORDER_ITEM, oi.OrderID, oi.ProductID, oi.Quantity).Scan(&oiID)
	if err != nil {
		return nil, err
	}
	return &models.OrderItem{ID: oiID, OrderID: oi.OrderID, ProductID: oi.ProductID, Quantity: oi.Quantity}, nil
}

// Updates an order item in the database.
func (tb *OrderItemsTable) UpdateOrderItem(ctx context.Context, oi *models.OrderItem) error {
	// validate data
	if err := utils.ValidateUpdateOrderItemData(oi); err != nil {
		return err
	}
	_, err := tb.DB.Exec(ctx, SQL_UPDATE_ORDER_ITEM, oi.OrderID, oi.ProductID, oi.Quantity, oi.ID)
	return err
}

// Deletes an order item from the database.
func (tb *OrderItemsTable) DeleteOrderItem(ctx context.Context, id int) error {
	_, err := tb.DB.Exec(ctx, SQL_DELETE_ORDER_ITEM, id)
	return err
}