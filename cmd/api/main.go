package main

import (
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/viniciusjsls/go-api/docs"
	"github.com/viniciusjsls/go-api/internal/handlers"
	"github.com/viniciusjsls/go-api/internal/middleware"
)

// @title           Go API
// @version         1.0
// @description     Learning Go API with Gin
// @host            localhost:8080
// @BasePath        /
func main() {
	// Get the port from environment variable or use default
	// can be provided when calling `go run`
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// router, redirect to respective handdler
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/health", handlers.HealthHandler)
	router.GET("/users", handlers.UsersHandler)
	router.POST("/users", handlers.CreateUser)

	router.Run(":" + port)

	router.Use(middleware.Logging)
}
