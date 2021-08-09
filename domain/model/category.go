package model

// Category 항목에 대한 모델
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
