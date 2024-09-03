package categories

type Category struct {
	ID          int    `json:"id"` // unique identifier
	Name        string `json:"name"` // name of the category
	Description string `json:"description"` // description of the category
	Offered     bool   `json:"offered"` // whether the category is offered
}

type Categories struct {
	Data []*Category `json:"data"`
}

type SubCategory struct {
	ID          int    `json:"id"` // unique identifier
	Name        string `json:"name"` // name of the sub-category
	Description string `json:"description"` // description of the sub-category
	CategoryId  int    `json:"category"` // category ID of the sub-category
	Offered     bool   `json:"offered"` // whether the sub-category is offered
}

type SubCategories struct {
	Data []*SubCategory `json:"data"`
}