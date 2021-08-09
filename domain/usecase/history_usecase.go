package usecase

import (
	"go-accountbook/domain/model"
	"go-accountbook/domain/repository"
)

// HistoryUsecase is interactor of History
type HistoryUsecase interface {
	RegisterHistory(*model.History) error
	GetHistories() ([]*model.History, error)
	GetHistory(id int) (*model.History, error)
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

func (u *historyUsecase) GetHistories() (histories []*model.History, err error) {
	histories, err = u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return
}

func (u *historyUsecase) GetHistory(id int) (history *model.History, err error) {
	history, err = u.repo.Find(id)
	return
}
