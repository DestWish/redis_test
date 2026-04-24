package service

import (
	"context"
	"fmt"

	"github.com/DestWish/redis_test/internal/models"
	"github.com/DestWish/redis_test/internal/repository"
	"github.com/redis/go-redis/v9"
)

type User_service struct {
	Repo *repository.UserRepository
}

func New_userService(repo *repository.UserRepository, cache *redis.Client) *User_service{
	return &User_service{Repo: repo}
}

func (s *User_service) CreateUser(ctx context.Context, reqUser *models.User ) uint {
	ID, err := s.Repo.Create(ctx, reqUser)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return ID
}