package handlers

import "net/http"

func IsValidHttpRequestMethod(w *http.ResponseWriter, r *http.Request, allowedMethod string) bool {
	if r.Method != allowedMethod {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		return false
	}

	return true
}
