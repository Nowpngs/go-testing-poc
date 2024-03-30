package main

import (
	"database/sql"
	server "go-testing-poc/internal/api/http/server"
	"log"
	"net/http"

	_ "go-testing-poc/docs"

	_ "github.com/lib/pq"
)

// @title 			Golang Testing POC API
// @version         1.0
// @description     This is a sample server celler server.

// @host      		localhost:8080
// @BasePath  		/api/v1
func main() {
	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/golang-testing?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Println("Successfully connected to the database")

	// Pass the database connection to your server
	router := server.NewRouter(db)

	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
