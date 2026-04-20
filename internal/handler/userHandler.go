package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DestWish/redis_test/internal/models"
	"github.com/DestWish/redis_test/internal/service"
)

type userHandler struct {
	service *service.User_service
}

func New_userHandler(service *service.User_service) *userHandler{
	return &userHandler{service: service}
}

func (h *userHandler) Create (req models.Create_userRequest) *models.User{
	user := models.User{Email: req.Email, Name: req.Name}
	resUser := h.service.CreateUser(&user)
	return resUser
}
func (h *userHandler) Get (ctx context.Context, ID int) *models.User {
	var cached models.User
	key := strconv.Itoa(ID)
	err := h.service.RedisClient.HGetAll(ctx, key).Scan(&cached)
	if err == nil && cached.ID != 0 {
        return &cached
    }
	fmt.Printf("Ошибка поиска в кэше %v \n", err)

	user, err := h.service.Repo.GetUser(uint(ID))
	if err != nil {
		fmt.Printf("Ошибка поиска в бд %v \n", err)
		return nil
	}
	fmt.Printf("Запись в кэш... \n")
	if err := h.service.RedisClient.HSet(ctx, key, user).Err(); err != nil {
		fmt.Printf("Ошибка кэширования: %v \n ответ из бд: \n", err)
		return &user
	}

	if err := h.service.RedisClient.HGetAll(ctx, key).Scan(&cached); err == nil && cached.ID != 0 {
		fmt.Println("ответ из кэша: ")
        return &cached
    }
	fmt.Println("ответ из бд(в кжше не нашлось, даже после записи): ")
	return &user
}