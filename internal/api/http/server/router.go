package server

import (
	"database/sql"
	"go-testing-poc/internal/api/http/user"
	"go-testing-poc/internal/config"
	"go-testing-poc/pkg/user/repository"
	"go-testing-poc/pkg/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(db *sql.DB, cfg *config.Config) *gin.Engine {
	gin.SetMode(cfg.GinMode)

	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Init Repositories
	userRepo := repository.NewUserRepository(db)

	// Init Services
	userService := service.NewUserService(userRepo)

	// Init Handlers
	userHandler := user.NewUserHandler(userService)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Routes
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/", userHandler.CreateUserHandler)
		userRoutes.GET("/", userHandler.GetUserListHandler)
		userRoutes.GET("/:id", userHandler.GetUserByIdHandler)
	}

	return router
}
