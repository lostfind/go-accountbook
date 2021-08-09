package repository

import (
	"go-accountbook/domain/model"
)

// HistoryRepository is Repository for History
type HistoryRepository interface {
	Save(history *model.History) error
	Find(id int) (*model.History, error)
	FindAll() ([]*model.History, error)
}
