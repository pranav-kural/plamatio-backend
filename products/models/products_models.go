
package products

// Product represents a product in the system.
type Product struct {
	ID             int    `json:"id"`             // unique identifier
	Name           string `json:"name"`           // name of the product
	Description    string `json:"description"`    // description of the product
	CategoryId     int    `json:"category"`       // category of the product
	SubCategoryId  int    `json:"subCategory"`    // sub-category of the product
	ImageURL       string `json:"imageUrl"`       // URL to the product image
	Price          int    `json:"price"`          // price of the product in cents
	PreviousPrice  int    `json:"previousPrice"`  // previous price of the product in cents
	Offered        bool   `json:"offered"`        // whether the product is offered
}

// Products represents a collection of products.
type Products struct {
	Data []*Product `json:"data"`
}

// ProductRequestParams represents the parameters required to create or update a product.
type ProductRequestParams struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	CategoryId     int    `json:"category"`
	SubCategoryId  int    `json:"subCategory"`
	ImageURL       string `json:"imageUrl"`
	Price          int    `json:"price"`
	PreviousPrice  int    `json:"previousPrice"`
	Offered        bool   `json:"offered"`
}

// ErrNameRequired is the error message for when the product name is missing.
const ErrNameRequired = "product name is required"

// ErrCategoryInvalid is the error message for when the product category is invalid.
const ErrCategoryInvalid = "invalid product category"

// ErrImageURLRequired is the error message for when the product image URL is missing.
const ErrImageURLRequired = "product image URL is required"

// ErrPriceInvalid is the error message for when the product price is invalid.
const ErrPriceInvalid = "product price must be greater than 0"