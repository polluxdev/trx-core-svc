package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/polluxdev/trx-core-svc/application/config"
)

type httpContext struct {
	c *gin.Context
}

type Meta struct {
	CurrentPage      int64 `json:"current_page"`
	LimitPerPage     int64 `json:"limit_per_page"`
	TotalCurrentPage int64 `json:"total_current_page"`
	TotalPages       int64 `json:"total_pages"`
	TotalData        int64 `json:"total_data"`
}

type Response struct {
	Code        int         `json:"code"`
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	MessageCode string      `json:"message_code"`
	Meta        *Meta       `json:"meta,omitempty"`
	Data        interface{} `json:"data"`
	Errors      interface{} `json:"errors,omitempty"`
}

func ToJSON(c *gin.Context) *httpContext {
	return &httpContext{c}
}

func (hc httpContext) CustomResponse(code int, success bool, msg string, msgCode string, data interface{}, errors interface{}) {
	hc.c.JSON(code, &Response{
		Code:        code,
		Success:     success,
		Message:     msg,
		MessageCode: msgCode,
		Data:        data,
		Errors:      errors,
	})
}

func (hc httpContext) PaginationResponse(code int, success bool, msg string, msgCode string, data interface{}, errors error, currentPage int64, limit int64, total int64) {
	totalPages := total / limit
	if total%limit > 0 {
		totalPages++
	}

	var totalCurrentPage int64
	if currentPage < totalPages {
		totalCurrentPage = limit
	} else if currentPage == totalPages {
		totalCurrentPage = total - (limit * (currentPage - 1))
	}

	response := &Response{
		Code:        code,
		Success:     success,
		Message:     msg,
		MessageCode: msgCode,
		Meta: &Meta{
			CurrentPage:      currentPage,
			LimitPerPage:     limit,
			TotalCurrentPage: totalCurrentPage,
			TotalPages:       totalPages,
			TotalData:        total,
		},
		Data: data,
	}

	if config.App.AppDebug && errors != nil {
		response.Errors = errors.Error()
	}

	hc.c.JSON(code, response)
}
