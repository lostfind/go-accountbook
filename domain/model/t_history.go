package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lostfind/go-accountbook/domain/app"
)

// THistory is t_history model
type THistory struct {
	gorm.Model
	AccountID  int
	CategoryID int
	Amount     int
	Memo       string
	Date       time.Time
	Balance    int
	// historyType  int `gorm:"column:type"` // いらないかも
	// linkid       int
	// creditCardID int
	// linkType int
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
