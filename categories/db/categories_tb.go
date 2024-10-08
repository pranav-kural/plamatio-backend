package categories

import (
	"context"

	models "encore.app/categories/models"
	utils "encore.app/categories/utils"
	"encore.dev/storage/sqldb"
)

type CategoriesTable struct {
	DB *sqldb.Database
}

/*

CREATE TABLE categories (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    offered BOOLEAN NOT NULL,
);

CREATE TABLE sub_categories (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    category_id BIGINT NOT NULL,
    offered BOOLEAN NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

*/

const (
    SQL_GET_CATEGORY = `
        SELECT name, description, offered FROM categories
        WHERE id = $1
    `
    SQL_GET_ALL_CATEGORIES = `
        SELECT id, name, description, offered FROM categories
    `
)

// Retrieves a category from the database.
func (tb *CategoriesTable) Get(ctx context.Context, id int) (*models.Category, error) {
	// validate id
	err := utils.ValidateProductCategory(id)
	if err != nil {
		return nil, err
	}
	c := &models.Category{ID: id}
	err = tb.DB.QueryRow(ctx, SQL_GET_CATEGORY, id).Scan(&c.Name, &c.Description, &c.Offered)
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
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.Offered); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return &models.Categories{Data: categories}, nil
}