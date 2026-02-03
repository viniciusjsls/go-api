package main

import (
	"log"
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
	mux.HandleFunc("/users", handlers.UsersHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("listening on :%s", port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
