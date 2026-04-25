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

func (s *User_service) ReadUser(ctx context.Context, req *models.ReadUserRequest) models.User {
	user, err := s.Repo.GetUser(ctx, req.ID)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return user
}

func (s *User_service) ReplaceUser(ctx context.Context, req *models.UpdateUserRequest) bool{
	user := models.User{ID: req.ID, Name: req.Name, Email: req.Email}
	ok, err := s.Repo.ReplaceUser(ctx, user)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return ok
	}

	return ok
}

func (s *User_service) PatchUser(ctx context.Context, req *models.PatchUserRequest) bool{
	user := models.User{ID: req.ID, Name: req.Name, Email: req.Email}
	ok, err := s.Repo.PatchUser(ctx, user)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return ok
	}

	return ok
}


func (s *User_service) DeleteUser(ctx context.Context, req *models.DeleteUserRequest) bool {
	ok, err := s.Repo.Delete(ctx, req.ID)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return ok
	}

	return ok
}
