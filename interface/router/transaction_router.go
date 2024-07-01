package router

import (
	"github.com/gin-gonic/gin"
	"github.com/polluxdev/trx-core-svc/interface/controller"
)

func SetupTransactionRouter(router *gin.RouterGroup, controller *controller.TransactionController) {
	r := router.Group("/transactions")
	r.POST("/checkout", controller.CreateTransaction)
}
