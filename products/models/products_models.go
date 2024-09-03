package products

/*

CREATE TABLE products (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    category_id BIGINT NOT NULL,
    sub_category_id BIGINT NOT NULL,
    image_url TEXT NOT NULL,
    price INT NOT NULL,
    previous_price INT,
    offered BOOLEAN NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
    FOREIGN KEY (sub_category_id) REFERENCES subcategories(id)
);

*/

type Product struct {
	ID       int    `json:"id"` // unique identifier
	Name     string `json:"name"` // name of the product
	Description string `json:"description"` // description of the product
	CategoryId  int `json:"category"`// category of the product
	SubCategoryId  int `json:"subCategory"`// sub-category of the product
	ImageURL string `json:"imageUrl"` // URL to the product image
	Price    int    `json:"price"` // price of the product in cents
	PreviousPrice int `json:"previousPrice"` // previous price of the product in cents
	Offered  bool   `json:"offered"` // whether the product is offered
}

type Products struct {
  Data []*Product `json:"data"`
}	

type ProductRequestParams struct {
	Name     string `json:"name"`
	Description string `json:"description"`
	CategoryId int `json:"category"`
	SubCategoryId int `json:"subCategory"`
	ImageURL string `json:"imageUrl"`
	Price    int    `json:"price"`
	PreviousPrice int `json:"previousPrice"`
	Offered  bool   `json:"offered"`
}

const ErrNameRequired = "product name is required"
const ErrCategoryInvalid = "invalid product category"
const ErrImageURLRequired = "product image URL is required"
const ErrPriceInvalid = "product price must be greater than 0"