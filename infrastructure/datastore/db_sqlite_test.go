package datastore

import (
	"testing"

	"go-accountbook/conf"

	"github.com/stretchr/testify/assert"
)

func TestNewSqliteDB(t *testing.T) {
	conf.Read()
	db, err := NewSqliteDB()

	assert.NoError(t, err)
	assert.NotNil(t, db)
}
