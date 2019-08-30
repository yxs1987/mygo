package model

type Category struct {
	Tm
	CategoryId    int64      `json:"category_id"`
	CategoryName  string     `json:"category_name"`
	ChildCategory []Category `json:"child_category"`
	Image         string     `json:"image"`
}
