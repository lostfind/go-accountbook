package usecase

import (
	"errors"
	"go-accountbook/domain/model"
	"go-accountbook/domain/repository"
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockHistoryRepository struct {
}

func NewMockHistoryRepository() *mockHistoryRepository {
	return &mockHistoryRepository{}
}

func (r *mockHistoryRepository) Save(history *model.History) error {
	// history.ID = rand.Intn(10)

	return nil
}

func (r *mockHistoryRepository) Find(id int) (history *model.History, err error) {
	history = new(model.History)

	history.ID = id
	history.AccountID = rand.Intn(10)
	history.CategoryID = rand.Intn(10)
	history.Amount = rand.Intn(50000)
	history.Memo = "TEST"
	history.Date = "2020/01/05"

	return
}

func (r *mockHistoryRepository) FindAll() (histories model.Histories, err error) {

	for i := 1; i < 10; i++ {
		history := &model.History{
			ID:         i,
			AccountID:  rand.Intn(10),
			CategoryID: rand.Intn(10),
			Amount:     rand.Intn(50000),
			Memo:       "TEST",
			Date:       "2020/01/05",
		}

		histories = append(histories, history)
	}

	return
}

func (r *mockHistoryRepository) FindAccounts(id []int) (map[int]*model.Account, error) {
	dummyAccounts := map[int]*model.Account{
		1: {ID: 1, Balance: 5000},
		2: {ID: 2, Balance: 5000},
		3: {ID: 3, Balance: 3000},
	}

	accounts := make(map[int]*model.Account)
	for _, i := range id {
		account, ok := dummyAccounts[i]
		if !ok {
			return accounts, errors.New("존재하지 않는 자산")
		}
		accounts[i] = account
	}

	return accounts, nil
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
		Date:       "2020/01/05",
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

func Test_historyUsecase_MoveAsset(t *testing.T) {
	historyRepository := NewMockHistoryRepository()
	type fields struct {
		repo repository.HistoryRepository
	}
	type args struct {
		amount        int
		fromAccountID int
		toAccountID   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Histories
		wantErr bool
	}{
		{
			name:   "자산이동 정상패턴",
			fields: fields{historyRepository},
			args:   args{amount: 5000, fromAccountID: 1, toAccountID: 2},
			want: model.Histories{
				{AccountID: 1, Amount: -5000},
				{AccountID: 2, Amount: 5000},
			},
		},
		{
			name:    "잔액이 이동금액보다 적은경우",
			fields:  fields{historyRepository},
			args:    args{amount: 5000, fromAccountID: 3, toAccountID: 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		// tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			u := &historyUsecase{
				repo: tt.fields.repo,
			}
			got, err := u.MoveAsset(tt.args.amount, tt.args.fromAccountID, tt.args.toAccountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("historyUsecase.MoveAsset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("historyUsecase.MoveAsset() = %v, want %v", got, tt.want)
			}
		})
	}
}
