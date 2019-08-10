package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewHistory(t *testing.T) {
	accountID := 5
	categoryID := 30
	amount := 1080
	memo := "Test Memo"
	date, _ := time.Parse("2006-01-02", "2019-08-10")

	history := NewHistory(accountID, categoryID, amount, memo, date)

	assert.Equal(t, accountID, history.AccountID)
	assert.Equal(t, categoryID, history.CategoryID)
	assert.Equal(t, amount, history.Amount)
	assert.Equal(t, memo, history.Memo)
	assert.Equal(t, date, history.Date)
}
