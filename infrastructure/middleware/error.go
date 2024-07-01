package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polluxdev/trx-core-svc/application/config"
	"github.com/polluxdev/trx-core-svc/common/utils"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if res, ok := err.(*utils.Response); ok {
					c.AbortWithStatusJSON(res.Code, res)
					return
				}

				if config.App.AppDebug {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"code":  http.StatusInternalServerError,
						"error": err,
					})
				} else {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"code":    http.StatusInternalServerError,
						"message": "Internal Server Error",
					})
				}
			}
		}()
		c.Next()
	}
}
