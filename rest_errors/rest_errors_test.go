package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("test error ISE", errors.New("database error"))

	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "test error ISE", err.Message())
	assert.EqualValues(t, "message: test error ISE - status: 500 - error: internal_server error - causes: [database error]", err.Error())

	assert.NotNil(t, err.Causes())
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("test error BRE")

	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "test error BRE", err.Message())
	assert.EqualValues(t, "message: test error BRE - status: 400 - error: bad_request - causes: []", err.Error())

	assert.Nil(t, err.Causes())
	assert.EqualValues(t, 0, len(err.Causes()))
}

func TestNewNotFoundError(t *testing.T) {
	err := NewBadRequestError("test error NFE")

	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "test error NFE", err.Message())
	assert.EqualValues(t, "message: test error NFE - status: 400 - error: bad_request - causes: []", err.Error())

	assert.Nil(t, err.Causes())
	assert.EqualValues(t, 0, len(err.Causes()))
}

func TestNewError(t *testing.T) {
	err := NewRestError("test error new", http.StatusBadRequest, "test error BRE", []interface{}{errors.New("new err")})

	assert.NotNil(t, err)
	assert.EqualValues(t, "message: test error new - status: 400 - error: test error BRE - causes: [new err]", err.Error())
}
