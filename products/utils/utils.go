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
		return errors.New("product name is required")
	}
	if p.Description == "" {
		return errors.New("product description is required")
	}
	if p.CategoryId < 0 || p.CategoryId > 3 {
		return errors.New("invalid product category; should be between 0 and 3")
	}
	if p.SubCategoryId < 0 || p.SubCategoryId > 10 {
		return errors.New("invalid product sub-category; should be between 0 and 10")
	}
	if p.ImageURL == "" {
		return errors.New("product image URL is required")
	}
	if p.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}

	return nil
}

