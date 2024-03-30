package server

import (
	"database/sql"
	"fmt"
	"net/http"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// Health check endpoint
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Server is running...")
    })

	return mux
}