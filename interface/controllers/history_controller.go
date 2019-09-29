package controllers

import (
	"github.com/lostfind/go-accountbook/entity/model"
	"github.com/lostfind/go-accountbook/usecase/interactor"
)

type historyController struct {
	historyInteractor interactor.HistoryUsecase
}

type HistoryController interface {
	GetHistories() ([]*model.History, error)
	RegisterHistory(history *model.History) error
}

// NewHistoryController is Constructor History controller
func NewHistoryController(i interactor.HistoryUsecase) HistoryController {
	return &historyController{
		historyInteractor: i,
	}
}

func (c *historyController) GetHistories() ([]*model.History, error) {
	return c.historyInteractor.GetHistories()
}

func (c *historyController) RegisterHistory(history *model.History) error {
	return c.historyInteractor.RegisterHistory(history)
}

func (c *historyController) GetHistory(id int) (history *model.History, err error) {
	return c.historyInteractor.GetHistory(id)
}
