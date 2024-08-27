package products

import (
	"context"

	models "encore.app/products/models"
	utils "encore.app/products/utils"
	"encore.dev/storage/sqldb"
)

type CategoriesTable struct {
	DB *sqldb.Database
}

const (
    SQL_GET_CATEGORY = `
        SELECT name, description, image_url, offered FROM categories
        WHERE id = $1
    `
    SQL_GET_ALL_CATEGORIES = `
        SELECT id, name, description, image_url, offered FROM categories
    `
)

// Retrieves a category from the database.
func (tb *CategoriesTable) Get(ctx context.Context, id int) (*models.Category, error) {
	// validate id
	err := utils.ValidateProductCategory(id)
	if err != nil {
		return nil, err
	}
	c := &models.Category{ID: models.ProductCategory(id)}
	err = tb.DB.QueryRow(ctx, SQL_GET_CATEGORY, id).Scan(&c.Name, &c.Description, &c.ImageURL, &c.Offered)
	return c, err
}

// Retrieves all categories from the database.
func (tb *CategoriesTable) GetAll(ctx context.Context) (*models.Categories, error) {
	rows, err := tb.DB.Query(ctx, SQL_GET_ALL_CATEGORIES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		c := &models.Category{}
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.ImageURL, &c.Offered); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return &models.Categories{Data: categories}, nil
}