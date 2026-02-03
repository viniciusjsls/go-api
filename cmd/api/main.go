package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/viniciusjsls/go-api/internal/handlers"
	"github.com/viniciusjsls/go-api/internal/middleware"
)

func main() {
	// Get the port from environment variable or use default
	// can be provided when calling `go run`
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// router, redirect to respective handdler
	router := gin.Default()
	router.GET("/health", handlers.HealthHandler)
	router.GET("/users", handlers.UsersHandler)
	router.POST("/users", handlers.CreateUser)

	router.Run(":" + port)

	router.Use(middleware.Logging)
}
