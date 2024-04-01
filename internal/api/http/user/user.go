package user

import (
	userModal "go-testing-poc/pkg/user"
	userDto "go-testing-poc/pkg/user/dto"
	userService "go-testing-poc/pkg/user/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService userService.UserService
}

func NewUserHandler(userService userService.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags Users
// @Accept  json
// @Param user body userDto.UserCreateRequest true "User object that needs to be created"
// @Produce  json
// @Success 201 {object} userModal.User
// @Router /user [post]
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var userRequest userDto.UserCreateRequest
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := userModal.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Role:     userRequest.Role,
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
// @Param role query string false "Role ID"
// @Produce  json
// @Tags Users
// @Success 200 {array} userModal.User
// @Router /user [get]
func (h *UserHandler) GetUserListHandler(c *gin.Context) {
	roleStr := c.Query("role")
	var role *userModal.Role

	if roleStr != "" {
		if roleID, err := strconv.Atoi(roleStr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
			return
		} else {
			role = (*userModal.Role)(&roleID)
		}
	}

	users, err := h.userService.GetUserList(c.Request.Context(), role)
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
// @Success 200 {object} userModal.User
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
