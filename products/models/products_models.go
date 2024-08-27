package products

type Product struct {
	ID       int    `json:"id"` // unique identifier
	Name     string `json:"name"` // name of the product
	Description string `json:"description"` // description of the product
	CategoryId ProductCategory `json:"category"`// category of the product
	ImageURL string `json:"imageUrl"` // URL to the product image
	Price    int    `json:"price"` // price of the product in cents
}

type Products struct {
  Data []*Product `json:"data"`
}	

type ProductRequestParams struct {
	Name     string `json:"name"`
	Description string `json:"description"`
	CategoryId int `json:"category"`
	ImageURL string `json:"imageUrl"`
	Price    int    `json:"price"`
}

const ErrNameRequired = "product name is required"
const ErrCategoryInvalid = "invalid product category ID: Must be a valid number. 1: Hats, 2: Dresses, 3: Tops, 4: Sneakers, 5: Jackets"
const ErrImageURLRequired = "product image URL is required"
const ErrPriceInvalid = "product price must be greater than 0"