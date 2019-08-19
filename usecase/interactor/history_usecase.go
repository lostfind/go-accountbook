package interactor

import (
	"github.com/lostfind/go-accountbook/entity/model"
	"github.com/lostfind/go-accountbook/usecase/repository"
)

// HistoryUsecase is interactor of History
type HistoryUsecase interface {
	RegisterHistory(*model.History) error
	ListHistory() ([]*model.History, error)
}

type historyUsecase struct {
	repo repository.HistoryRepository
}

// NewHistoryUsecase return HistoryUsecase
func NewHistoryUsecase(repo repository.HistoryRepository) HistoryUsecase {
	return &historyUsecase{
		repo: repo,
	}
}

func (u *historyUsecase) RegisterHistory(history *model.History) error {
	err := u.repo.Save(history)
	if err != nil {
		return err
	}
	return nil
}

func (u *historyUsecase) ListHistory() ([]*model.History, error) {
	histories, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return histories, nil
}
