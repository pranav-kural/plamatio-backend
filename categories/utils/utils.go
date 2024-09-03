package categories

import (
	"errors"
)

func ValidateProductCategory(categoryId int) error {
	// validate the product category
	if categoryId < 0 || categoryId > 3 {
		return errors.New("invalid category id; can be from 1 to 3")
	}

	return nil
}

func ValidateProductSubCategory(subCategoryId int) error {
	// validate the product sub-category
	if subCategoryId < 0 || subCategoryId > 10 {
		return errors.New("invalid sub-category id; can be from 1 to 10")
	}

	return nil
}

