package service

import (
	"fmt"

	"github.com/DestWish/redis_test/internal/models"
	"github.com/DestWish/redis_test/internal/repository"
	"github.com/redis/go-redis/v9"
)

type User_service struct {
	Repo *repository.UserRepository
	RedisClient *redis.Client
}

func New_userService(repo *repository.UserRepository, cache *redis.Client) *User_service{
	return &User_service{Repo: repo, RedisClient: cache}
}

func (s *User_service) CreateUser(reqUser *models.User ) *models.User {
	user, err := s.Repo.Create(reqUser)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return user
}