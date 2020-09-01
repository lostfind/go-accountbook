package datastore

import (
	"testing"

	"github.com/lostfind/go-accountbook/conf"
	"github.com/stretchr/testify/assert"
)

func TestNewMySQLDB(t *testing.T) {
	conf.Read()
	db, err := NewMySQLDB()

	assert.NoError(t, err)
	assert.NotNil(t, db)
}
