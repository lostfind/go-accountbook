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
	accountType := 1

	account := NewAccount(name, balance, accountType)

	assert.Equal(t, name, account.Name)
	assert.Equal(t, balance, account.Balance)
	assert.Equal(t, accountType, account.AccountType)
}
