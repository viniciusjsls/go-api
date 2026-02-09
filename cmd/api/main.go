package main

import (
	"log"

	_ "github.com/viniciusjsls/go-api/docs"
	"github.com/viniciusjsls/go-api/internal/app"
	"github.com/viniciusjsls/go-api/internal/config"
)

// @title           Go API
// @version         1.0
// @description     Learning Go API with Gin
// @host            localhost:8080
// @BasePath        /
func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	logger := config.InitLogger(cfg)

	server := app.NewServer(cfg, logger)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
