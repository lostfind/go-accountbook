package datastore

import (
	"testing"

	"go-accountbook/conf"

	"github.com/stretchr/testify/assert"
)

func TestNewMySQLDB(t *testing.T) {
	conf.Read()
	db, err := NewMySQLDB()

	assert.NoError(t, err)
	assert.NotNil(t, db)
}
