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
	if cartItem.UserID <= 0 {
		return errors.New("invalid user ID")
	}
	return nil
}