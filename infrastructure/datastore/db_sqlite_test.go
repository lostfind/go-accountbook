package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSqliteDB(t *testing.T) {
	db := NewSqliteDB()

	assert.NotNil(t, db)
}
