package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/application/libs"
	"github.com/polluxdev/trx-core-svc/application/model"
	"github.com/polluxdev/trx-core-svc/application/service"
	"github.com/polluxdev/trx-core-svc/common/helper"
	"github.com/polluxdev/trx-core-svc/common/utils"
	"github.com/polluxdev/trx-core-svc/interface/serializer"
)

type LimitController struct {
	service   *service.LimitService
	validator *libs.CustomValidator
}

func NewLimitController(service *service.LimitService, validator *libs.CustomValidator) *LimitController {
	return &LimitController{service: service, validator: validator}
}

func (c *LimitController) GetLimitList(ctx *gin.Context) {
	query := new(model.PaginationQuery)
	response := utils.ToJSON(ctx)

	if err := ctx.ShouldBindQuery(query); err != nil {
		response.CustomResponse(http.StatusBadRequest, false, global.BAD_REQUEST, "BAD_REQUEST_QUERY_STRING", nil, err)
		return
	}

	query.Page = helper.SetDefaultIfZero(query.Page, 1)
	query.Limit = helper.SetDefaultIfZero(query.Limit, 10)

	result, total := c.service.GetLimitList(query)

	response.PaginationResponse(
		http.StatusOK,
		true,
		global.GET_DATA_SUCCESS,
		"GET_LIMIT_LIST_SUCCESS",
		serializer.SerializeLimits(result),
		nil,
		int64(query.Page),
		int64(query.Limit),
		*total,
	)
}
