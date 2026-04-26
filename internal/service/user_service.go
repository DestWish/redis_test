package service

import (
	"context"

	"github.com/DestWish/redis_test/internal/models"
	"github.com/DestWish/redis_test/internal/repository"
)

type User_service struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *User_service {
	return &User_service{Repo: repo}
}

func (s *User_service) CreateUser(ctx context.Context, req *models.CreateUserRequest) (uint, error) {
	return s.Repo.Create(ctx, req)
}

func (s *User_service) ReadUser(ctx context.Context, req *models.ReadUserRequest) (models.User, error) {
	return s.Repo.GetUser(ctx, req)
}

func (s *User_service) ReplaceUser(ctx context.Context, req *models.UpdateUserRequest) (bool, error) {
	return s.Repo.ReplaceUser(ctx, req)
}

func (s *User_service) PatchUser(ctx context.Context, req *models.PatchUserRequest) (bool, error) {
	return s.Repo.PatchUser(ctx, req)
}

func (s *User_service) DeleteUser(ctx context.Context, req *models.DeleteUserRequest) (bool, error) {
	return s.Repo.Delete(ctx, req)
}
