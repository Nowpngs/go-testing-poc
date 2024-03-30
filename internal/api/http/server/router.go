package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"go-testing-poc/internal/api/http/user"
	"go-testing-poc/pkg/user/repository"
	"go-testing-poc/pkg/user/service"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running...")
	})

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.CreateUserHandler(w, r)
		case http.MethodGet:
			userHandler.GetUserListHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
