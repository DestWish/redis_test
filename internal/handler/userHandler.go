package handler

import (
	"context"
	"fmt"

	"github.com/DestWish/redis_test/internal/models"
	"github.com/DestWish/redis_test/internal/service"
)

type userHandler struct {
	service *service.User_service
}

func New_userHandler(service *service.User_service) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Create(ctx context.Context, req models.Create_userRequest) uint {
	user := models.User{Email: req.Email, Name: req.Name}
	ID := h.service.CreateUser(ctx, &user)
	return ID
}

func (h *userHandler) Get(ctx context.Context, ID uint) *models.User {
	user, err := h.service.Repo.GetUser(ctx, ID)
	if err != nil {
		fmt.Printf("Ошибка запроса: %v", nil)
		return nil
	}
	return &user
}
