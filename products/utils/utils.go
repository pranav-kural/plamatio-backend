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

func ValidateProductRequestParams(p *models.ProductRequestParams) error {
	// validate the product request parameters
	if p.Name == "" {
		return errors.New(models.ErrNameRequired)
	}
	if p.CategoryId < 0 || p.CategoryId > 5 {
		return errors.New(models.ErrCategoryInvalid)
	}
	if p.ImageURL == "" {
		return errors.New(models.ErrImageURLRequired)
	}
	if p.Price <= 0 {
		return errors.New(models.ErrPriceInvalid)
	}

	return nil
}

