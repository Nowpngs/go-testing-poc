package server

import (
	"database/sql"
	"go-testing-poc/internal/api/http/booking"
	"go-testing-poc/internal/api/http/user"
	"go-testing-poc/internal/config"
	bookingRepo "go-testing-poc/pkg/booking/repository"
	bookingService "go-testing-poc/pkg/booking/service"
	userRepo "go-testing-poc/pkg/user/repository"
	userService "go-testing-poc/pkg/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(db *sql.DB, cfg *config.Config) *gin.Engine {
	gin.SetMode(cfg.GinMode)

	router := gin.Default()

	router.GET("/docs/*any", func(c *gin.Context) {
		if c.Request.URL.Path == "/docs" {
			c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
			return
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})

	// Init Repositories
	userRepo := userRepo.NewUserRepository(db)
	bookingRepo := bookingRepo.NewBookingRepository(db)

	// Init Services
	userService := userService.NewUserService(userRepo)
	bookingService := bookingService.NewBookingService(bookingRepo)

	// Init Handlers
	userHandler := user.NewUserHandler(userService)
	bookingHandler := booking.NewBookingHandler(bookingService)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Routes
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("", userHandler.CreateUserHandler)
		userRoutes.GET("", userHandler.GetUserListHandler)
		userRoutes.GET("/:id", userHandler.GetUserByIdHandler)
	}
	bookingRoutes := router.Group("/booking")
	{
		bookingRoutes.POST("", bookingHandler.CreateBookingHandler)
		bookingRoutes.GET("user", bookingHandler.GetBookingWithUserListHandler)
	}

	return router
}
