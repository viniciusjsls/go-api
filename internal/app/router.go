package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/viniciusjsls/go-api/internal/config"
	"github.com/viniciusjsls/go-api/internal/handlers"
	"github.com/viniciusjsls/go-api/internal/middleware"
	"go.uber.org/zap"
)

func InitRouter(cfg *config.Config, logger *zap.Logger) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/health", handlers.HealthHandler)
	router.GET("/users", handlers.UsersHandler)
	router.POST("/users", handlers.CreateUser)

	router.Use(func(c *gin.Context) {
		middleware.Logging(c, logger)
	})

	return router
}
