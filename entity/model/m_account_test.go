package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBalance(t *testing.T) {
	ma := new(MAccount)
	ma.SetBalance(100)

	assert.Equal(t, 100, ma.Balance)
}
