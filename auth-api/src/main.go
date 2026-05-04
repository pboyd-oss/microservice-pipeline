package main

import (
	"log"
	"net/http"

	"auth-api/src/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health)
	mux.HandleFunc("POST /auth/login", handlers.Login)
	mux.HandleFunc("POST /auth/validate", handlers.Validate)

	log.Println("auth-api listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
