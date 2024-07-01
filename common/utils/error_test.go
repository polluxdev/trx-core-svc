package utils

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/polluxdev/trx-core-svc/application/config"
	"github.com/stretchr/testify/assert"
)

func TestClientError(t *testing.T) {
	err := errors.New("client error")
	expected := &Response{
		Code:        http.StatusBadRequest,
		Message:     "client error message",
		MessageCode: "BAD_REQUEST",
		Errors:      "client error",
	}

	config.App.AppDebug = true
	response := ClientError("client error message", err)
	assert.Equal(t, expected, response, "Expected response to be equal")

	config.App.AppDebug = false
	expected.Errors = nil
	response = ClientError("client error message", err)
	assert.Equal(t, expected, response, "Expected response to be equal when debug mode is off")
}

func TestInvariantError(t *testing.T) {
	err := errors.New("invariant error")
	expected := &Response{
		Code:        http.StatusInternalServerError,
		Message:     "invariant error message",
		MessageCode: "INTERNAL_SERVER_ERROR",
		Errors:      "invariant error",
	}

	config.App.AppDebug = true
	response := InvariantError("invariant error message", err)
	assert.Equal(t, expected, response, "Expected response to be equal")

	config.App.AppDebug = false
	expected.Message = "Internal Server Error"
	expected.Errors = nil
	response = InvariantError("invariant error message", err)
	assert.Equal(t, expected, response, "Expected response to be equal when debug mode is off")
}

func TestNotFoundError(t *testing.T) {
	err := errors.New("not found error")
	expected := &Response{
		Code:        http.StatusNotFound,
		Message:     "not found error message",
		MessageCode: "DATA_NOT_FOUND",
		Errors:      "not found error",
	}

	config.App.AppDebug = true
	response := NotFoundError("not found error message", err)
	assert.Equal(t, expected, response, "Expected response to be equal")

	config.App.AppDebug = false
	expected.Errors = nil
	response = NotFoundError("not found error message", err)
	assert.Equal(t, expected, response, "Expected response to be equal when debug mode is off")
}

func TestIsNotFoundError(t *testing.T) {
	notFoundErr := gorm.ErrRecordNotFound
	otherErr := errors.New("some other error")

	assert.True(t, IsNotFoundError(notFoundErr), "Expected true for gorm.ErrRecordNotFound")
	assert.False(t, IsNotFoundError(otherErr), "Expected false for other error")
}
