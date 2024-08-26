package products

import (
	"errors"

	models "encore.app/products/models"
)

func ValidateProductCategory(category string) error {
	// confirm the category is valid
	if category != string(models.Hats) && category != string(models.Dresses) && category != string(models.Tops) && category != string(models.Sneakers) && category != string(models.Jackets) {
		return errors.New(models.ErrCategoryInvalid)
	}
	return nil
}

func ValidateProductRequestParams(p *models.ProductRequestParams) error {
	// validate the product request parameters
	if p.Name == "" {
		return errors.New(models.ErrNameRequired)
	}
	if p.Category == "" {
		return errors.New(models.ErrCategoryRequired)
	}
	if p.ImageURL == "" {
		return errors.New(models.ErrImageURLRequired)
	}
	if p.Price <= 0 {
		return errors.New(models.ErrPriceInvalid)
	}

	// confirm the category is valid
	if err := ValidateProductCategory(p.Category); err != nil {
		return err
	}

	return nil
}

