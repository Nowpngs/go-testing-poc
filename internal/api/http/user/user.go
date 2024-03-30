package user

import (
	modal "go-testing-poc/pkg/user"
	"go-testing-poc/pkg/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags Users
// @Accept  json
// @Param user body modal.User true "User object that needs to be created"
// @Produce  json
// @Success 201 {object} modal.User
// @Router /user [post]
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var newUser modal.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.CreateUser(c.Request.Context(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

// GetUserList godoc
// @Summary Get user list
// @Description Get a list of users
// @Produce  json
// @Tags Users
// @Success 200 {array} modal.User
// @Router /user [get]
func (h *UserHandler) GetUserListHandler(c *gin.Context) {
	users, err := h.userService.GetUserList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserById godoc
// @Summary Get user by ID
// @Description Get a single user by ID
// @Param id path string true "User ID"
// @Produce  json
// @Tags Users
// @Success 200 {object} modal.User
// @Router /user/{id} [get]
func (h *UserHandler) GetUserByIdHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUserById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
