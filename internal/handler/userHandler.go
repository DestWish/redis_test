package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DestWish/redis_test/internal/models"
	"github.com/DestWish/redis_test/internal/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service *service.User_service
}

func NewUserHandler(service *service.User_service) *userHandler {
	return &userHandler{service: service}
}

func (h * userHandler) RegisterRoutes(r *gin.Engine){
	users := r.Group("api/users")
	{
		users.POST("", h.Create)
		users.GET("")
		users.PUT("")
		users.DELETE("")
	}
}

func (h *userHandler) Create(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	userID:= h.service.CreateUser(ctx, &req)

	c.JSON(http.StatusCreated, userID)
}

func (h *userHandler) Read(ctx context.Context, ID uint) *models.User {
	user, err := h.service.Repo.GetUser(ctx, ID)
	if err != nil {
		fmt.Printf("Ошибка запроса: %v", err)
		return nil
	}
	return &user
}
