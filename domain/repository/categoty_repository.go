package repository

import "go-accountbook/domain/model"

type CategoryRepository interface {
	Category(id int) (*model.Category, error)
	NewCategory(category *model.Category) error
	AllCategoies() (model.Categories, error)
}
