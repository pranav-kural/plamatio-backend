package products

type Product struct {
	ID       int    // unique identifier
	Name     string // name of the product
	ImageURL string // URL to the product image
	Price    int    // price of the product in cents
}