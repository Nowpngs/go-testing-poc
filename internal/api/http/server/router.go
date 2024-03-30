package server

import (
	"database/sql"
	"go-testing-poc/internal/api/http/user"
	"go-testing-poc/pkg/user/repository"
	"go-testing-poc/pkg/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/", userHandler.CreateUserHandler)
		userRoutes.GET("/", userHandler.GetUserListHandler)
		userRoutes.GET("/:id", userHandler.GetUserByIdHandler)
	}

	return router
}
