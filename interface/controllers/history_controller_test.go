package controllers

import (
	"math/rand"
	"testing"
	"time"

	"go-accountbook/domain/model"
	"go-accountbook/domain/usecase"

	"github.com/stretchr/testify/assert"
)

type mockHistoryUsecase struct {
}

func NewMockHistoryUsecase() usecase.HistoryUsecase {
	return &mockHistoryUsecase{}
}

func (u *mockHistoryUsecase) RegisterHistory(history *model.History) (err error) {
	history.ID = rand.Intn(10)
	return
}

func (u *mockHistoryUsecase) GetHistories() (histories []*model.History, err error) {
	count := rand.Intn(10)

	for i := 0; i < count; i++ {
		history := new(model.History)
		histories = append(histories, history)
	}
	return
}

func (u *mockHistoryUsecase) GetHistory(id int) (history *model.History, err error) {
	return
}

func TestGetHistories(t *testing.T) {
	usecase := NewMockHistoryUsecase()
	controller := NewHistoryController(usecase)

	histories, err := controller.GetHistories()

	assert.NoError(t, err)
	assert.NotEmpty(t, histories)
}

func TestRegisterHistory(t *testing.T) {
	usecase := NewMockHistoryUsecase()
	controller := NewHistoryController(usecase)

	history := model.NewHistory(1, 1, 4000, "TEST", time.Now())
	err := controller.RegisterHistory(history)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, history.ID)
}
