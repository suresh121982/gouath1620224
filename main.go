package main

import (
	"log"
	"net/http"

	"examples.com/restapi-cache-auth/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Authentication routes
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")


	// Example protected route with authorization middleware
	r.Handle("/cache", handlers.AuthMiddleware(http.HandlerFunc(handlers.CacheHandler))).Methods("POST")
	r.Handle("/example", handlers.AuthMiddleware(http.HandlerFunc(handlers.ExampleHandler))).Methods("GET")

	// Initialize cache
	handlers.InitializeCache()

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
