package model

// Category 항목에 대한 모델
type Category struct {
	ID         int
	Name       string
	Count      int
	Sort       int
	DeleteFlag bool
}

type Categories []*Category

// NewCategory is constructor Category
func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}
