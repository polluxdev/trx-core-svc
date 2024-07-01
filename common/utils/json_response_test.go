package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/polluxdev/trx-core-svc/application/config"
	"github.com/stretchr/testify/assert"
)

func TestCustomResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	hc := ToJSON(c)
	hc.CustomResponse(http.StatusOK, true, "success message", "SUCCESS_CODE", "data", nil)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{"code":200,"success":true,"message":"success message","message_code":"SUCCESS_CODE","data":"data"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestPaginationResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	hc := ToJSON(c)
	hc.PaginationResponse(http.StatusOK, true, "pagination message", "PAGINATION_CODE", "paged data", nil, 1, 10, 50)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{
		"code": 200,
		"success": true,
		"message": "pagination message",
		"message_code": "PAGINATION_CODE",
		"meta": {
			"current_page": 1,
			"limit_per_page": 10,
			"total_current_page": 10,
			"total_pages": 5,
			"total_data": 50
		},
		"data": "paged data"
	}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestPaginationResponseWithErrors(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	config.App.AppDebug = true // Enable debug mode for this test

	hc := ToJSON(c)
	hc.PaginationResponse(http.StatusOK, true, "pagination message with error", "PAGINATION_CODE", "paged data", assert.AnError, 1, 10, 50)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{
		"code": 200,
		"success": true,
		"message": "pagination message with error",
		"message_code": "PAGINATION_CODE",
		"meta": {
			"current_page": 1,
			"limit_per_page": 10,
			"total_current_page": 10,
			"total_pages": 5,
			"total_data": 50
		},
		"data": "paged data",
		"errors": "assert.AnError general error for testing"
	}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestPaginationResponseLastPage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	hc := ToJSON(c)
	hc.PaginationResponse(http.StatusOK, true, "pagination message last page", "PAGINATION_CODE", "paged data", nil, 5, 10, 50)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{
		"code": 200,
		"success": true,
		"message": "pagination message last page",
		"message_code": "PAGINATION_CODE",
		"meta": {
			"current_page": 5,
			"limit_per_page": 10,
			"total_current_page": 10,
			"total_pages": 5,
			"total_data": 50
		},
		"data": "paged data"
	}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
