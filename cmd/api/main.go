package main

import (
	"net/http"
	"os"

	"github.com/viniciusjsls/go-api/internal/handlers"
)

func main() {
	// Get the port from environment variable or use default
	// can be provided when calling `go run`
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// router, redirect to respective handdler
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
