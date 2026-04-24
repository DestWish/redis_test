package repository

import (
	"context"
	"fmt"

	"github.com/DestWish/redis_test/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewUserRepo(db *gorm.DB, redisClient *redis.Client) *UserRepository {
	return &UserRepository{db: db, redisClient: redisClient}
}

func (r *UserRepository) userCaching(ctx context.Context, user *models.User) error {
	key := userCacheKey(user.ID)
	if err := r.redisClient.HSet(ctx, key, user).Err(); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (uint, error) {
	result := r.db.Create(user)
	if err := result.Error; err != nil {
		fmt.Printf("Error %v", err)
		return user.ID, err
	}

	if err := r.userCaching(ctx, user); err != nil {
		return user.ID, err
	}

	return user.ID, nil
}

func (r *UserRepository) GetUser(ctx context.Context, userID uint) (models.User, error) {
	key := userCacheKey(userID)
	var user models.User

	err := r.redisClient.HGetAll(ctx, key).Scan(&user)
	if err == nil && user.ID != 0 {
		return user, nil
	}

	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		fmt.Printf("Данного пользователя не существует %v\n", err.Error())
		return user, err
	}

	if err := r.userCaching(ctx, &user); err != nil {
		fmt.Printf("Ошибка кэширования %v\n", err.Error())
	}

	return user, nil
}

func userCacheKey(userID uint) string {
	return fmt.Sprintf("user:%v", userID)
}
