package cart

import (
	"errors"

	models "encore.app/cart/models"
)

func ValidateCartData(cartItem *models.CartItem, idRequired bool, allowEmpty bool) error {
	if !allowEmpty && cartItem == nil {
		return errors.New("empty cart item object")
	}
	if idRequired && cartItem.ID <= 0 {
		return errors.New("invalid cart item ID")
	}
	if cartItem.ProductID <= 0 {
		return errors.New("invalid product ID")
	}
	if cartItem.Quantity <= 0 {
		return errors.New("invalid quantity, cannot be less than 1")
	}
	if cartItem.UserID == "" {
		return errors.New("invalid user ID")
	}
	return nil
}

func ValidateNewCartItem(newCartItem *models.NewCartItem) error {
	if newCartItem == nil {
		return errors.New("empty new cart item object")
	}
	if newCartItem.ProductID <= 0 {
		return errors.New("invalid product ID")
	}
	if newCartItem.Quantity <= 0 {
		return errors.New("invalid quantity, cannot be less than 1")
	}
	if newCartItem.UserID == "" {
		return errors.New("invalid user ID")
	}
	return nil
}

func ValidateNewCartItems(newCartItems *models.NewCartItems) error {
	if newCartItems == nil {
		return errors.New("empty new cart items object")
	}
	if len(newCartItems.Data) == 0 {
		return errors.New("empty new cart items list")
	}
	for _, newCartItem := range newCartItems.Data {
		if err := ValidateNewCartItem(newCartItem); err != nil {
			return err
		}
	}
	return nil
}