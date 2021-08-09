package model

import (
	"time"
)

// History 이력 모델
type History struct {
	ID         int       `gorm:"primary_key"`
	AccountID  int       `gorm:"account_id"`
	CategoryID int       `gorm:"category_id"`
	Amount     int       `gorm:"amount"`
	Memo       string    `gorm:"memo"`
	Date       time.Time `gorm:"date"`
}
