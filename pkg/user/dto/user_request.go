package dto

import (
	userModal "go-testing-poc/pkg/user"
)

type UserCreateRequest struct {
	Username string         `json:"username" binding:"required"`
	Email    *string        `json:"email"`
	Role     userModal.Role `json:"role" binding:"required"`
}
