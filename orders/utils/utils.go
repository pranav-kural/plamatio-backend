package orders

import (
	"errors"
	"time"

	models "encore.app/orders/models"
)

func ValidateNewOrderData(data *models.OrderRequestParams) error {
	if data.UserID <= 0 {
		return errors.New("user_id is required")
	}
	if data.AddressID <= 0 {
		return errors.New("address_id is required")
	}
	if data.TotalPrice <= 0 {
		return errors.New("total_price is required")
	}
	if data.Status == "" {
		return errors.New("status is required")
	}
	// attempt to parse the created_at string
	// should follow the RFC3339 format
	_, err := time.Parse(time.RFC3339, data.CreatedAt)
	if err != nil {
		return errors.New("created_at is required and should follow the RFC3339 format")
	}
	return nil
}

func ValidateUpdateOrderData(data *models.Order) error {
	if data.ID <= 0 {
		return errors.New("id is required")
	}
	return ValidateNewOrderData(&models.OrderRequestParams{
		UserID: data.UserID,
		AddressID: data.AddressID,
		TotalPrice: data.TotalPrice,
		CreatedAt: data.CreatedAt,
		Status: data.Status,
	})
}

func ValidateNewOrderItemData(data *models.OrderItemRequestParams) error {
	if data.OrderID <= 0 {
		return errors.New("order_id is required")
	}
	if data.ProductID <= 0 {
		return errors.New("product_id is required")
	}
	if data.Quantity <= 0 {
		return errors.New("quantity is required")
	}
	return nil
}

func ValidateUpdateOrderItemData(data *models.OrderItem) error {
	if data.ID <= 0 {
		return errors.New("id is required")
	}
	return ValidateNewOrderItemData(&models.OrderItemRequestParams{
		OrderID: data.OrderID,
		ProductID: data.ProductID,
		Quantity: data.Quantity,
	})
}