package controllers

import (
	"go-accountbook/domain/model"
	"go-accountbook/domain/usecase"
)

type historyController struct {
	usecase usecase.HistoryUsecase
}

type HistoryController interface {
	GetHistories() ([]*model.History, error)
	RegisterHistory(history *model.History) error
}

// NewHistoryController is Constructor History controller
func NewHistoryController(i usecase.HistoryUsecase) HistoryController {
	return &historyController{
		usecase: i,
	}
}

func (c *historyController) GetHistories() ([]*model.History, error) {
	return c.usecase.GetHistories()
}

func (c *historyController) RegisterHistory(history *model.History) error {
	return c.usecase.RegisterHistory(history)
}

func (c *historyController) GetHistory(id int) (history *model.History, err error) {
	return c.usecase.GetHistory(id)
}
