package products

import (
	"context"

	models "encore.app/products/models"
	utils "encore.app/products/utils"
	"encore.dev/storage/sqldb"
)

type ProductsTB struct {
	DB *sqldb.Database
}

const (
    SQL_INSERT_PRODUCT = `
        INSERT INTO products (name, description, category_id, sub_category_id, image_url, price, previous_price, offered)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id
    `
    SQL_DELETE_PRODUCT = `
        DELETE FROM products
        WHERE id = $1
    `
    SQL_UPDATE_PRODUCT = `
        UPDATE products
        SET name = $1, description = $2, category_id = $3, sub_category_id = $4, image_url = $5, price = $6, previous_price = $7, offered = $8
        WHERE id = $9
    `
		SQL_GET_PRODUCT = `
        SELECT name, description, category_id, sub_category_id, image_url, price, previous_price, offered FROM products
        WHERE id = $1
    `
    SQL_GET_ALL_PRODUCTS = `
        SELECT id, name, description, category_id, sub_category_id, image_url, price, previous_price, offered FROM products
    `
		SQL_GET_PRODUCTS_BY_CATEGORY = `
				SELECT id, name, description, category_id, sub_category_id, image_url, price, previous_price, offered FROM products
				WHERE category_id = $1
		`
		SQL_GET_PRODUCTS_BY_SUB_CATEGORY = `
				SELECT id, name, description, category_id, sub_category_id, image_url, price, previous_price, offered FROM products
				WHERE sub_category_id = $1
		`
		SQL_GET_HERO_PRODUCTS = `
				SELECT p.id, p.name, p.description, p.category_id, p.sub_category_id, p.image_url, p.price, p.previous_price, p.offered
				FROM products p
				INNER JOIN hero_products hp ON p.id = hp.product_id
		`
		SQL_GET_CATEGORY_HERO_PRODUCTS_BY_CATEGORY = `
				SELECT p.id, p.name, p.description, p.category_id, p.sub_category_id, p.image_url, p.price, p.previous_price, p.offered
				FROM products p
				INNER JOIN category_hero_products chp ON p.id = chp.product_id
				WHERE chp.category_id = $1
		`
)

// Inserts a product into the database.
// Inserts a product into the database and returns the id of the newly added record.
func (pdb *ProductsTB) Insert(ctx context.Context, p *models.ProductRequestParams) (int, error) {
    var id int
    err := pdb.DB.QueryRow(ctx, SQL_INSERT_PRODUCT, p.Name, p.Description, p.CategoryId, p.SubCategoryId, p.ImageURL, p.Price, p.PreviousPrice, p.Offered).Scan(&id)
		
    return id, err
}

// Bulk inserts products into the database.
func (pdb *ProductsTB) BulkInsert(ctx context.Context, products []*models.ProductRequestParams) error {
	tx, err := pdb.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := pdb.DB.Stdlib().Prepare(SQL_INSERT_PRODUCT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range products {
		// validate data
		err := utils.ValidateProductRequestParams(p)
		if err != nil {
			return err
		}
		if _, err := stmt.Exec(p.Name, p.Description, p.CategoryId, p.SubCategoryId, p.ImageURL, p.Price, p.PreviousPrice, p.Offered); err != nil {
			return err
		}
	}
  return nil
}

// Deletes a product from the database.
func (pdb *ProductsTB) Delete(ctx context.Context, id int) error {
	_, err := pdb.DB.Exec(ctx, SQL_DELETE_PRODUCT, id)
	return err
}

// Updates a product in the database.
func (pdb *ProductsTB) Update(ctx context.Context, id int, p *models.ProductRequestParams) error {
	_, err := pdb.DB.Exec(ctx, SQL_UPDATE_PRODUCT, p.Name, p.Description, p.CategoryId, p.SubCategoryId, p.ImageURL, p.Price, p.PreviousPrice, p.Offered, id)
	return err
}

// Retrieves a product from the database.
func (pdb *ProductsTB) Get(ctx context.Context, id int) (*models.Product, error) {
	p := &models.Product{ID: id}
	err := pdb.DB.QueryRow(ctx, SQL_GET_PRODUCT, id).Scan(&p.Name, &p.Description, &p.CategoryId, &p.SubCategoryId, &p.ImageURL, &p.Price, &p.PreviousPrice, &p.Offered)
	return p, err
}

// Retrieves all products from the database.
func (pdb *ProductsTB) GetAll(ctx context.Context) (*models.Products, error) {
	rows, err := pdb.DB.Query(ctx, SQL_GET_ALL_PRODUCTS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CategoryId, &p.SubCategoryId, &p.ImageURL, &p.Price, &p.PreviousPrice, &p.Offered); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &models.Products{Data: products}, nil
}

// Retrieves all products from the database by category.
func (pdb *ProductsTB) GetByCategory(ctx context.Context, id int) (*models.Products, error) {
	rows, err := pdb.DB.Query(ctx, SQL_GET_PRODUCTS_BY_CATEGORY, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CategoryId, &p.SubCategoryId, &p.ImageURL, &p.Price, &p.PreviousPrice, &p.Offered); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &models.Products{Data: products}, nil
}

// Retrieves all products from the database by sub-category.
func (pdb *ProductsTB) GetBySubCategory(ctx context.Context, id int) (*models.Products, error) {
	rows, err := pdb.DB.Query(ctx, SQL_GET_PRODUCTS_BY_SUB_CATEGORY, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CategoryId, &p.SubCategoryId, &p.ImageURL, &p.Price, &p.PreviousPrice, &p.Offered); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &models.Products{Data: products}, nil
}

// Retrieve hero products from the database.
func (pdb *ProductsTB) GetHeroProducts(ctx context.Context) (*models.Products, error) {
	// Query the database for hero products.
	rows, err := pdb.DB.Query(ctx, SQL_GET_HERO_PRODUCTS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and create a product for each row.
	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CategoryId, &p.SubCategoryId, &p.ImageURL, &p.Price, &p.PreviousPrice, &p.Offered); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &models.Products{Data: products}, nil
}

// Retrieve hero products from the database by category.
func (pdb *ProductsTB) GetCategoryHeroProducts(ctx context.Context, id int) (*models.Products, error) {
	// Query the database for hero products by category.
	rows, err := pdb.DB.Query(ctx, SQL_GET_CATEGORY_HERO_PRODUCTS_BY_CATEGORY, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and create a product for each row.
	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CategoryId, &p.SubCategoryId, &p.ImageURL, &p.Price, &p.PreviousPrice, &p.Offered); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &models.Products{Data: products}, nil
}

// Retrieve products based on search query.
func (pdb *ProductsTB) Search(ctx context.Context, query string) (*models.Products, error) {
	// Query the database for products based on search query.
	rows, err := pdb.DB.Query(ctx, "SELECT id, name, description, category_id, image_url FROM products WHERE name ILIKE $1", "%"+query+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and create a product for each row.
	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CategoryId, &p.SubCategoryId, &p.ImageURL, &p.Price, &p.PreviousPrice, &p.Offered); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return &models.Products{Data: products}, nil
}