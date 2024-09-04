package orders

import (
	"context"

	models "encore.app/orders/models"
	utils "encore.app/orders/utils"
	"encore.dev/storage/sqldb"
)

type OrdersTable struct {
	DB *sqldb.Database
}

const (
		SQL_GET_ORDER = `
				SELECT user_id, address_id, total_price, created_at, status FROM orders
				WHERE id = $1
		`
		SQL_GET_ALL_ORDERS = `
				SELECT id, user_id, address_id, total_price, created_at, status FROM orders
		`
		SQL_GET_ORDERS_BY_USER = `
				SELECT id, user_id, address_id, total_price, created_at, status FROM orders
				WHERE user_id = $1
		`
		SQL_INSERT_ORDER = `
				INSERT INTO orders (user_id, address_id, total_price, created_at, status) VALUES ($1, $2, $3, $4, $5) RETURNING id
		`
		SQL_UPDATE_ORDER = `
				UPDATE orders SET user_id = $1, address_id = $2, total_price = $3, created_at = $4, status = $5 WHERE id = $6
		`
		SQL_DELETE_ORDER = `
				DELETE FROM orders WHERE id = $1
		`
)

// Retrieves an order from the database.
func (tb *OrdersTable) GetOrder(ctx context.Context, id int) (*models.Order, error) {
	o := &models.Order{ID: id}
	err := tb.DB.QueryRow(ctx, SQL_GET_ORDER, id).Scan(&o.UserID, &o.AddressID, &o.TotalPrice, &o.CreatedAt, &o.Status)
	return o, err
}

// Retrieves all orders from the database.
func (tb *OrdersTable) GetAllOrders(ctx context.Context) (*models.Orders, error) {
	rows, err := tb.DB.Query(ctx, SQL_GET_ALL_ORDERS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := &models.Orders{}
	for rows.Next() {
		o := &models.Order{}
		if err := rows.Scan(&o.ID, &o.UserID, &o.AddressID, &o.TotalPrice, &o.CreatedAt, &o.Status); err != nil {
			return nil, err
		}
		orders.Data = append(orders.Data, o)
	}
	return orders, nil
}

// Retrieves all orders for a user from the database.
func (tb *OrdersTable) GetOrdersByUser(ctx context.Context, userId int) (*models.Orders, error) {
	rows, err := tb.DB.Query(ctx, SQL_GET_ORDERS_BY_USER, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := &models.Orders{}
	for rows.Next() {
		o := &models.Order{}
		if err := rows.Scan(&o.ID, &o.UserID, &o.AddressID, &o.TotalPrice, &o.CreatedAt, &o.Status); err != nil {
			return nil, err
		}
		orders.Data = append(orders.Data, o)
	}
	return orders, nil
}

// Inserts an order into the database.
func (tb *OrdersTable) InsertOrder(ctx context.Context, o *models.OrderRequestParams) (*models.Order, error) {
	// validate data
	if err := utils.ValidateNewOrderData(o); err != nil {
		return nil, err
	}
	// insert order
	var id int
	err := tb.DB.QueryRow(ctx, SQL_INSERT_ORDER, o.UserID, o.AddressID, o.TotalPrice, o.CreatedAt, o.Status).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &models.Order{ID: id, UserID: o.UserID, AddressID: o.AddressID, TotalPrice: o.TotalPrice, CreatedAt: o.CreatedAt, Status: o.Status}, nil
}

// Updates an order in the database.
func (tb *OrdersTable) UpdateOrder(ctx context.Context, o *models.Order) error {
	// validate data
	if err := utils.ValidateUpdateOrderData(o); err != nil {
		return err
	}
	_, err := tb.DB.Exec(ctx, SQL_UPDATE_ORDER, o.UserID, o.AddressID, o.TotalPrice, o.CreatedAt, o.Status, o.ID)
	return err
}

// Deletes an order from the database.
func (tb *OrdersTable) DeleteOrder(ctx context.Context, id int) error {
	_, err := tb.DB.Exec(ctx, SQL_DELETE_ORDER, id)
	return err
}