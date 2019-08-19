package controllers

import (
	"github.com/lostfind/go-accountbook/entity/model"
	"github.com/lostfind/go-accountbook/usecase/interactor"
)

type historyController struct {
	historyInteractor interactor.HistoryUsecase
}

type HistoryController interface {
	GetListHistory() ([]*model.History, error)
	RegisterHistory(history *model.History) error
}

func NewHistoryController(i interactor.HistoryUsecase) HistoryController {
	return &historyController{
		historyInteractor: i,
	}
}

func (c *historyController) GetListHistory() ([]*model.History, error) {
	return c.historyInteractor.ListHistory()
}

func (c *historyController) RegisterHistory(history *model.History) error {
	return c.historyInteractor.RegisterHistory(history)
}
