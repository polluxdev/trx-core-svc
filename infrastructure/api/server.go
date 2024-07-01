package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polluxdev/trx-core-svc/application/config"
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/application/libs"
	"github.com/polluxdev/trx-core-svc/application/service"
	"github.com/polluxdev/trx-core-svc/infrastructure/database"
	"github.com/polluxdev/trx-core-svc/infrastructure/middleware"
	"github.com/polluxdev/trx-core-svc/infrastructure/repository"
	"github.com/polluxdev/trx-core-svc/interface/controller"
	"github.com/polluxdev/trx-core-svc/interface/router"
	"github.com/sirupsen/logrus"
)

func Serve() {
	if config.App.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(gin.LoggerWithWriter(logrus.StandardLogger().Writer()))

	logrus.SetFormatter(&logrus.JSONFormatter{})

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"host":    c.Request.Host,
			"message": config.App.AppName,
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"host":    c.Request.Host,
			"path":    c.Request.URL.Path,
			"message": "Page not found",
		})
	})

	apiV1 := r.Group("/api/v1")

	// Initialize libs
	validator := libs.NewCustomValidator()

	// Initialize database
	dbConn := database.NewConnection()

	// Initialize repositories
	consumerRepository := repository.NewConsumerRepository(dbConn.Db)
	limitRepository := repository.NewLimitRepository(dbConn.Db)

	// Initialize services
	consumerService := service.NewConsumerService(consumerRepository)
	limitService := service.NewLimitService(limitRepository)

	// Initialize controllers
	consumerController := controller.NewConsumerController(consumerService, validator)
	limitController := controller.NewLimitController(limitService, validator)

	// Setup routers
	router.SetupConsumerRouter(apiV1, consumerController)
	router.SetupLimitRouter(apiV1, limitController)

	if err := r.Run(":" + config.App.AppPort); err != nil {
		logrus.WithFields(logrus.Fields{
			"app_version": global.BUILD_VERSION,
			"app_port":    config.App.AppPort,
			"error":       err.Error(),
		})
	}
}
