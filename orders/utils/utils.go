package orders

import (
	"errors"

	models "encore.app/orders/models"
)

func ValidateNewOrderData(data *models.OrderRequestParams) error {
	if data.UserID == "" {
		return errors.New("user_id is required")
	}
	if data.AddressID <= 0 {
		return errors.New("address_id is required")
	}
	if data.TotalPrice <= 0.0 {
		return errors.New("total_price is required")
	}
	if data.Status == "" {
		return errors.New("status is required")
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