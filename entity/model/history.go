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
}

// NewHistory is constructor History
func NewHistory(account, category, amount int, memo string, date time.Time) *History {
	return &History{
		AccountID:  account,
		CategoryID: category,
		Amount:     amount,
		Memo:       memo,
		Date:       date,
	}
}

// TableName return Table Name
func (History) TableName() string {
	return "t_history"
}

// GetHistory return public fields for app.History
func (t *History) GetHistory() *app.History {
	return &app.History{
		Amount: t.Amount,
		Memo:   t.Memo,
	}
}
