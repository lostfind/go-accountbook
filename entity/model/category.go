package model

// Category is m_categoris model
type Category struct {
	ID         int    `gorm:"primary_key"`
	Name       string `gorm:"name"`
	Count      int    `gorm:"count"`
	Sort       int    `gorm:"sort"`
	DeleteFlag bool   `gorm:"is_delete"`
}

// NewCategory is constructor Category
func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}

// TableName return Table Name
func (Category) TableName() string {
	return "m_categories"
}
