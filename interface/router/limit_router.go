package router

import (
	"github.com/gin-gonic/gin"
	"github.com/polluxdev/trx-core-svc/interface/controller"
)

func SetupLimitRouter(router *gin.RouterGroup, controller *controller.LimitController) {
	r := router.Group("/limits")
	r.GET("", controller.GetLimitList)
}
