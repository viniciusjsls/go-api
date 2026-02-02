package main

import (
	"os"
)

func main() {
	// Get the port from environment variable or use default
	// can be provided when calling `go run`
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
}
