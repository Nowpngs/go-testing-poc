package service

import (
	"context"
	userModal "go-testing-poc/pkg/user"
	userRepo "go-testing-poc/pkg/user/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *userModal.User) error
	GetUserList(ctx context.Context, role *userModal.Role) ([]*userModal.User, error)
	GetUserById(ctx context.Context, id string) (*userModal.User, error)
}

type userService struct {
	userRepo userRepo.UserRepository
}

func NewUserService(userRepo userRepo.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, user *userModal.User) error {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) GetUserList(ctx context.Context, role *userModal.Role) ([]*userModal.User, error) {
	return s.userRepo.GetUserList(ctx, role)
}

func (s *userService) GetUserById(ctx context.Context, id string) (*userModal.User, error) {
	return s.userRepo.GetUserById(ctx, id)
}
