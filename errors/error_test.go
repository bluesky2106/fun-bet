package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	assert := assert.New(t)

	err := ErrMySQLConnErrtion
	err = Wrap(err, "Msg 1", "Msg 2")

	assert.NotNil(err)
	assert.Equal("Msg 2: Msg 1: mysql connection error", err.Error(), "error message mismatched")
}
