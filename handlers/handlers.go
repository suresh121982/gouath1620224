package handlers

import (
	"net/http"
)

// InitializeCache initializes the cache map
func InitializeCache() {
	// Initialize cache map
	cache = make(map[string][]byte)
}

// Define your handlers here

// Other handlers related to authentication, etc.

// NotFoundHandler handles 404 errors
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Not Found"))
}
