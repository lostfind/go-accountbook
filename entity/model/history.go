package model

import (
	"time"
)

// History is t_histories model
type History struct {
	ID         int       `gorm:"primary_key"`
	AccountID  int       `gorm:"account_id"`
	CategoryID int       `gorm:"category_id"`
	Amount     int       `gorm:"amount"`
	Memo       string    `gorm:"memo"`
	Date       time.Time `gorm:"date"`
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
	return "t_histories"
}
