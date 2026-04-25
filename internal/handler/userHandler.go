package handler

import (
	"net/http"
	"strconv"

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

func (h *userHandler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("api/users")
	{
		users.POST("", h.Create)
		users.GET("/:id", h.Read)
		users.PUT("", h.Update)
		users.PATCH("")
		users.DELETE("/:id", h.Delete)
	}
}

func (h *userHandler) Create(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	userID := h.service.CreateUser(ctx, &req)

	c.JSON(http.StatusCreated, userID)
}

func (h *userHandler) Read(c *gin.Context) {
	id, ok := parseUserId(c)
	if !ok {
		return
	}

	req := models.ReadUserRequest{ID: id}

	ctx := c.Request.Context()
	user := h.service.ReadUser(ctx, &req)

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) Update(c *gin.Context) {
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	success := h.service.ReplaceUser(ctx, &req)

	c.JSON(http.StatusOK, success)
}

func (h *userHandler) Delete(c *gin.Context) {
	id, ok := parseUserId(c)
	if !ok {
		return
	}

	req := models.DeleteUserRequest{ID: id}

	ctx := c.Request.Context()
	success := h.service.DeleteUser(ctx, &req)

	c.JSON(http.StatusNoContent, success)
}

func parseUserId(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return 0, false
	}

	return uint(id), true
}
