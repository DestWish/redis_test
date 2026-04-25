package service

import (
	"context"
	"fmt"

	"github.com/DestWish/redis_test/internal/models"
	"github.com/DestWish/redis_test/internal/repository"
)

type User_service struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *User_service {
	return &User_service{Repo: repo}
}

func (s *User_service) CreateUser(ctx context.Context, req *models.CreateUserRequest) uint {
	user := models.User{
		Name: req.Name,
		Email: req.Email,
	}

	ID, err := s.Repo.Create(ctx, &user)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return ID
}
