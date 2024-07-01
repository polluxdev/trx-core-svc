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

type ConsumerController struct {
	service   *service.ConsumerService
	validator *libs.CustomValidator
}

func NewConsumerController(service *service.ConsumerService, validator *libs.CustomValidator) *ConsumerController {
	return &ConsumerController{service: service, validator: validator}
}

func (c *ConsumerController) CreateConsumer(ctx *gin.Context) {
	request := new(model.NewConsumer)
	response := utils.ToJSON(ctx)

	if err := ctx.ShouldBindJSON(request); err != nil {
		errMsg := c.validator.ParseError(err)
		response.CustomResponse(http.StatusBadRequest, false, global.BAD_REQUEST, "BAD_REQUEST_BODY", nil, errMsg)
		return
	}

	result := c.service.CreateConsumer(request)

	response.CustomResponse(http.StatusCreated, true, global.CREATE_DATA_SUCCESS, "CREATE_CONSUMER_SUCCESS", serializer.SerializeConsumer(result), nil)
}
