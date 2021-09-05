package usecase

import (
	"errors"
	"go-accountbook/domain/model"
	"go-accountbook/domain/repository"

	log "github.com/sirupsen/logrus"
)

// HistoryUsecase is interactor of History
type HistoryUsecase interface {
	RegisterHistory(*model.History) error
	GetHistories() ([]*model.History, error)
	GetHistory(id int) (*model.History, error)
	MoveAsset(amount, fromAccountID, toAccountID int) (model.Histories, error)
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
		log.Error(err)
		return err
	}
	return nil
}

func (u *historyUsecase) GetHistories() (histories []*model.History, err error) {
	histories, err = u.repo.FindAll()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (u *historyUsecase) GetHistory(id int) (history *model.History, err error) {
	history, err = u.repo.Find(id)
	return
}

func (u *historyUsecase) MoveAsset(amount, fromAccountID, toAccountID int) (model.Histories, error) {
	accounts, err := u.repo.FindAccounts([]int{fromAccountID, toAccountID})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	fromAccount := accounts[fromAccountID]
	toAccount := accounts[toAccountID]
	fromAccount.UpdateBalance(amount * -1)
	toAccount.UpdateBalance(amount)

	if fromAccount.Balance < 0 {
		return nil, errors.New("잔액이 부족합니다")
	}

	fromHistory := model.History{}
	toHistory := model.History{}

	if err := u.repo.Save(&fromHistory); err != nil {
		log.Error(err)
		return nil, err
	}
	if err := u.repo.Save(&toHistory); err != nil {
		log.Error(err)
		return nil, err
	}
	histories := model.Histories{&fromHistory, &toHistory}

	return histories, nil
}
