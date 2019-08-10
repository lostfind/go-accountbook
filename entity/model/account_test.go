package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBalance(t *testing.T) {
	ma := new(Account)
	ma.SetBalance(100)

	assert.Equal(t, 100, ma.Balance)
}

func TestNewAccount(t *testing.T) {
	name := "TestAccount"
	balance := 1238523

	account := NewAccount(name, balance)

	assert.Equal(t, name, account.AccountName)
	assert.Equal(t, balance, account.Balance)
}
