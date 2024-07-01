package utils

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/polluxdev/trx-core-svc/application/config"
)

func ClientError(msg string, err error) *Response {
	response := &Response{
		Code:        http.StatusBadRequest,
		Message:     msg,
		MessageCode: "BAD_REQUEST",
	}

	if config.App.AppDebug && err != nil {
		response.Errors = err.Error()
	}

	return response
}

func InvariantError(msg string, err error) *Response {
	if IsNotFoundError(err) {
		return NotFoundError(msg, err)
	}

	response := &Response{
		Code:        http.StatusInternalServerError,
		MessageCode: "INTERNAL_SERVER_ERROR",
	}

	if config.App.AppDebug && err != nil {
		response.Message = msg
		response.Errors = err.Error()
	} else {
		response.Message = "Internal Server Error"
	}

	return response
}

func NotFoundError(msg string, err error) *Response {
	response := &Response{
		Code:        http.StatusNotFound,
		Message:     msg,
		MessageCode: "DATA_NOT_FOUND",
	}

	if config.App.AppDebug && err != nil {
		response.Errors = err.Error()
	}

	return response
}

func IsNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}
