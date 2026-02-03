package handlers

import (
	"encoding/json"
	"net/http"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if !IsValidHttpRequestMethod(&w, r, http.MethodGet) {
		return
	}

	users := []user{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(users)
}
