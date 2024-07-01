package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/common/utils"
)

func ValidateID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		if _, err := strconv.Atoi(idParam); err != nil {
			response := utils.ToJSON(c)
			response.CustomResponse(http.StatusBadRequest, false, global.BAD_REQUEST, "BAD_REQUEST_PARAMS", "ID must be a number", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
