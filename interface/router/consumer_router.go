package router

import (
	"github.com/gin-gonic/gin"
	"github.com/polluxdev/trx-core-svc/interface/controller"
)

func SetupConsumerRouter(router *gin.RouterGroup, controller *controller.ConsumerController) {
	r := router.Group("/consumers")
	r.POST("", controller.CreateConsumer)
}
