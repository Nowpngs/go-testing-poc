package service

import (
	"context"
	modal "go-testing-poc/pkg/user"
	"go-testing-poc/pkg/user/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *modal.User) error
	GetUserList(ctx context.Context) ([]*modal.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, user *modal.User) error {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) GetUserList(ctx context.Context) ([]*modal.User, error) {
	return s.userRepo.GetUserList(ctx)
}