package model

import (
	"time"

	"github.com/lostfind/go-accountbook/entity/app"
)

// History is t_history model
type History struct {
	AccountID  int
	CategoryID int
	Amount     int
	Memo       string
	Date       time.Time
	Balance    int
}

// NewHistory is constructor History
func NewHistory() *History {
	return &History{}
}

// TableName return Table Name
func (History) TableName() string {
	return "t_history"
}

// SetBalance calculate and set balance
func (t *History) SetBalance(preBalance int) {
	if t.Amount != 0 {
		t.Balance = preBalance + t.Amount
	}
}

// GetHistory return public fields for app.History
func (t *History) GetHistory() *app.History {
	return &app.History{
		Amount:  t.Amount,
		Balance: t.Balance,
		Memo:    t.Memo,
	}
}
