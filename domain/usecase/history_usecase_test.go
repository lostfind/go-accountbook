package usecase

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go-accountbook/domain/model"
	"go-accountbook/domain/repository"
)

type mockHistoryRepository struct {
}

func NewMockHistoryRepository() repository.HistoryRepository {
	return &mockHistoryRepository{}
}

func (r *mockHistoryRepository) Save(history *model.History) error {
	history.ID = rand.Intn(10)

	return nil
}

func (r *mockHistoryRepository) Find(id int) (history *model.History, err error) {
	history = new(model.History)

	history.ID = id
	history.AccountID = rand.Intn(10)
	history.CategoryID = rand.Intn(10)
	history.Amount = rand.Intn(50000)
	history.Memo = "TEST"
	history.Date = time.Now()

	return
}

func (r *mockHistoryRepository) FindAll() (histories []*model.History, err error) {
	return
}

func TestRegisterHistory(t *testing.T) {
	historyRepository := NewMockHistoryRepository()
	historyUsecase := NewHistoryUsecase(historyRepository)

	accountID := rand.Intn(10)
	categoryID := rand.Intn(10)
	amount := rand.Intn(50000)
	memo := "TEST"
	history := &model.History{
		AccountID:  accountID,
		CategoryID: categoryID,
		Amount:     amount,
		Memo:       memo,
		Date:       time.Now(),
	}

	historyUsecase.RegisterHistory(history)

	assert.NotEmpty(t, history.ID)
	assert.Equal(t, accountID, history.AccountID)
	assert.Equal(t, categoryID, history.CategoryID)
	assert.Equal(t, amount, history.Amount)
	assert.Equal(t, "TEST", history.Memo)
}

func TestGetHistory(t *testing.T) {
	historyRepository := NewMockHistoryRepository()
	historyUsecase := NewHistoryUsecase(historyRepository)

	testID := 15

	history, err := historyUsecase.GetHistory(testID)

	assert.NoError(t, err)
	assert.Equal(t, testID, history.ID)
}

func TestGetHistories(t *testing.T) {
	historyRepository := NewMockHistoryRepository()
	historyUsecase := NewHistoryUsecase(historyRepository)

	historyUsecase.GetHistories()
}
