package service

import (
	"context"

	"github.com/DestWish/redis_test/internal/models"
	"github.com/DestWish/redis_test/internal/repository"
)

type User_service struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *User_service {
	return &User_service{repo: repo}
}

func (s *User_service) CreateUser(ctx context.Context, req *models.CreateUserRequest) (string, error) {
	return s.repo.CreateUser(ctx, req)
}

func (s *User_service) ReadUser(ctx context.Context, req *models.ReadUserRequest) (models.User, error) {
	return s.repo.GetUser(ctx, req)
}

func (s *User_service) ReplaceUser(ctx context.Context, req *models.UpdateUserRequest) (bool, error) {
	return s.repo.ReplaceUser(ctx, req)
}

func (s *User_service) PatchUser(ctx context.Context, req *models.PatchUserRequest) (bool, error) {
	return s.repo.PatchUser(ctx, req)
}

func (s *User_service) DeleteUser(ctx context.Context, req *models.DeleteUserRequest) (bool, error) {
	return s.repo.DeleteUser(ctx, req)
}
