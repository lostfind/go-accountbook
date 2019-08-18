package repository

import "github.com/lostfind/go-accountbook/entity/model"

// HistoryRepository is Repository for History
type HistoryRepository interface {
	Save(*model.History) error
	FindAll() ([]*model.History, error)
}
