package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/application/libs"
	"github.com/polluxdev/trx-core-svc/application/model"
	"github.com/polluxdev/trx-core-svc/application/service"
	"github.com/polluxdev/trx-core-svc/common/utils"
	"github.com/polluxdev/trx-core-svc/interface/serializer"
)

type TransactionController struct {
	service   *service.TransactionService
	validator *libs.CustomValidator
}

func NewTransactionController(service *service.TransactionService, validator *libs.CustomValidator) *TransactionController {
	return &TransactionController{service: service, validator: validator}
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	request := new(model.NewTransaction)
	response := utils.ToJSON(ctx)

	if err := ctx.ShouldBindJSON(request); err != nil {
		errMsg := c.validator.ParseError(err)
		response.CustomResponse(http.StatusBadRequest, false, global.BAD_REQUEST, "BAD_REQUEST_BODY", nil, errMsg)
		return
	}

	result := c.service.CreateTransaction(request)

	response.CustomResponse(http.StatusCreated, true, global.CREATE_DATA_SUCCESS, "CREATE_TRANSACTION_SUCCESS", serializer.SerializeTransaction(result), nil)
}
