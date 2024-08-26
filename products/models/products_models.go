package products

type ProductCategory string

// Hats, Dresses, Tops, Sneakers, Jackets
const (
	Hats     ProductCategory = "Hats"
	Dresses  ProductCategory = "Dresses"
	Tops     ProductCategory = "Tops"
	Sneakers ProductCategory = "Sneakers"
	Jackets  ProductCategory = "Jackets"
)

type Product struct {
	ID       int    `json:"id"` // unique identifier
	Name     string `json:"name"` // name of the product
	Description string `json:"description"` // description of the product
	Category ProductCategory `json:"category"`// category of the product
	ImageURL string `json:"imageUrl"` // URL to the product image
	Price    int    `json:"price"` // price of the product in cents
}

type Products struct {
  Data []*Product `json:"data"`
}	

type ProductRequestParams struct {
	Name     string `json:"name"`
	Description string `json:"description"`
	Category string `json:"category"`
	ImageURL string `json:"imageUrl"`
	Price    int    `json:"price"`
}

const ErrNameRequired = "product name is required"
const ErrCategoryRequired = "product category is required"
const ErrCategoryInvalid = "invalid product category. must be one of: Hats, Dresses, Tops, Sneakers, Jackets"
const ErrImageURLRequired = "product image URL is required"
const ErrPriceInvalid = "product price must be greater than 0"