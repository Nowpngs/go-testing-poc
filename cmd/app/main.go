package main

import (
	"database/sql"
	"log"
	"net/http"

	server "go-testing-poc/internal/api/http/server"

	_ "github.com/lib/pq"
)

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
