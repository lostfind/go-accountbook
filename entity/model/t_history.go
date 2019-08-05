package model

import (
	"lostfind/go-accountbook/entity/app"
	"time"
)

// THistory is t_history model
type THistory struct {
	AccountID  int
	CategoryID int
	Amount     int
	Memo       string
	Date       time.Time
	Balance    int
}

// TableName is return Table Name
func (THistory) TableName() string {
	return "t_history"
}

// SetBalance is Calulate and set
func (t *THistory) SetBalance(preBalance int) {
	if t.Amount != 0 {
		t.Balance = preBalance + t.Amount
	}
}

// GetHistory is Return public fields for app.THistory
func (t *THistory) GetHistory() *app.THistory {
	return &app.THistory{
		Amount:  t.Amount,
		Balance: t.Balance,
		Memo:    t.Memo,
	}
}
