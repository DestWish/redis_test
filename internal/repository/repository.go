package repository

import (
	"context"

	"github.com/DestWish/redis_test/internal/models"
)

type UserRepo interface {
	CreateUser(context.Context, *models.CreateUserRequest) (uint, error) 
	GetUser(context.Context, *models.ReadUserRequest) (models.User, error)
	ReplaceUser(context.Context, *models.UpdateUserRequest) (bool, error)
	PatchUser(context.Context, *models.PatchUserRequest) (bool, error)
	DeleteUser(context.Context, *models.DeleteUserRequest) (bool, error)
}