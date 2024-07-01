package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Accept",
			"Accept-Encoding",
			"Authorization",
			"Cache-Control",
			"Content-Length",
			"Content-Type",
			"Origin",
			"X-CSRF-Token",
			"X-Requested-With",
			"X-XSRF-Token",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
