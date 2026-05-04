package main

import (
	"log"
	"net/http"

	"data-api/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health)
	mux.HandleFunc("GET /data", handlers.List)
	mux.HandleFunc("GET /data/{id}", handlers.Get)

	log.Println("data-api listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
