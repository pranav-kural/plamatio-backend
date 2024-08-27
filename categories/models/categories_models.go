package categories

type ProductCategory int

// Hats, Dresses, Tops, Sneakers, Jackets
const (
	Hats     ProductCategory = 1
	Dresses  ProductCategory = 2
	Tops     ProductCategory = 3
	Sneakers ProductCategory = 4
	Jackets  ProductCategory = 5
)

type Category struct {
	ID          ProductCategory    `json:"id"` // unique identifier
	Name        string `json:"name"` // name of the category
	Description string `json:"description"` // description of the category
	ImageURL    string `json:"imageUrl"` // URL to the category image
	Offered     bool   `json:"offered"` // whether the category is offered
}

type Categories struct {
	Data []*Category `json:"data"`
}