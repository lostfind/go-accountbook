package sqlrepo

import "go-accountbook/domain/model"

// Category
func (r *sqlRepo) Category(id int) (*model.Category, error) {
	return &model.Category{}, nil
}
func (r *sqlRepo) NewCategory(category *model.Category) error {
	return nil
}
func (r *sqlRepo) AllCategoies() (model.Categories, error) {
	return model.Categories{}, nil
}
