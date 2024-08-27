package products

import (
	"errors"

	models "encore.app/products/models"
)

func ValidateProductCategory(categoryId int) error {
	// validate the product category
	if categoryId < 0 || categoryId > 5 {
		return errors.New(models.ErrCategoryInvalid)
	}

	return nil
}

