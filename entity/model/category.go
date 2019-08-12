package model

// Category is m_category model
type Category struct {
	ID         int
	Name       string
	Count      int
	Sort       int
	DeleteFlag bool
}

// TableName return Table Name
func (Category) TableName() string {
	return "m_category"
}

// NewCategory is constructor Category
func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}
