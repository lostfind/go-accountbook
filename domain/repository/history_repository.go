package repository

import (
	"go-accountbook/domain/model"
)

// HistoryRepository is Repository for History
type HistoryRepository interface {
	Save(history *model.History) error
	Find(id int) (*model.History, error)
	FindAll() (model.Histories, error)
	FindAccounts(id []int) (map[int]*model.Account, error)
}
