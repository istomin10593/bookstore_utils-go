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

	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "test error ISE", err.Message)
	assert.EqualValues(t, "internal_server error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("test error BRE", errors.New("invalid err"))

	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "test error BRE", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "invalid err", err.Causes[0])
}

func TestNewNotFoundError(t *testing.T) {
	err := NewBadRequestError("test error NFE", errors.New("not found err"))

	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "test error NFE", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "not found err", err.Causes[0])
}

func TestNewError(t *testing.T) {
	err := NewError("test error new")

	assert.NotNil(t, err)
	assert.EqualValues(t, "test error new", err.Error())
}
