package server

import (
	"database/sql"
	"go-testing-poc/internal/api/http/user"
	"go-testing-poc/pkg/user/repository"
	"go-testing-poc/pkg/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
