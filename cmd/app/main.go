package main

import (
	"go-testing-poc/internal/api/http/server"
	"log"
	"net/http"
)

func main() {
    router := server.NewRouter()

    log.Println("Starting server on port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Server error: %v", err)
    }
}