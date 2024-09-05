package categories

// Category represents a category.
type Category struct {
	ID          int    `json:"id"`          // unique identifier
	Name        string `json:"name"`        // name of the category
	Description string `json:"description"` // description of the category
	Offered     bool   `json:"offered"`     // whether the category is offered
}

// Categories represents a collection of categories.
type Categories struct {
	Data []*Category `json:"data"`
}

// SubCategory represents a sub-category.
type SubCategory struct {
	ID          int    `json:"id"`          // unique identifier
	Name        string `json:"name"`        // name of the sub-category
	Description string `json:"description"` // description of the sub-category
	CategoryId  int    `json:"category"`    // category ID of the sub-category
	Offered     bool   `json:"offered"`     // whether the sub-category is offered
}

// SubCategories represents a collection of sub-categories.
type SubCategories struct {
	Data []*SubCategory `json:"data"`
}