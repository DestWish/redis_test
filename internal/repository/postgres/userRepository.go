package postgres

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
	if err := r.redisClient.HSet(ctx, user.Login, user).Err(); err != nil {
		return fmt.Errorf("Repo: Cache failed: %w", err)
	}

	return nil
}

func (r *UserRepository) CreateUser(ctx context.Context, req *models.CreateUserRequest) (string, error) {
	user := &models.User{Name: req.Name, Email: req.Email}
	if err := r.db.Model(&models.User{}).Create(user).Error; err != nil {
		return "", fmt.Errorf("Repo: Create user failed: %w", err)
	}

	return user.Login, r.userCaching(ctx, user)
}

func (r *UserRepository) GetUser(ctx context.Context, req *models.ReadUserRequest) (models.User, error) {
	var user models.User

	err := r.redisClient.HGetAll(ctx, req.Login).Scan(&user)
	if err == nil && req.Login != "" {
		return user, nil
	}

	if err := r.db.Where("id = ?", req.Login).First(&user).Error; err != nil {
		return user, fmt.Errorf("Repo: User not found: %w", err)
	}

	return user, r.userCaching(ctx, &user)
}

func (r *UserRepository) ReplaceUser(ctx context.Context, req *models.UpdateUserRequest) (bool, error) {
	result := r.db.Model(&models.User{}).Where("id = ?", req.Login).Select("*").Updates(req)

	if result.Error != nil {
		return false, fmt.Errorf("Repo: User not found: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return false, fmt.Errorf("Repo: User not found: %w", result.Error)
	}

	var updatedUser models.User
	if err := r.db.Where("id = ?", req.Login).First(&updatedUser).Error; err != nil {
		return false, fmt.Errorf("Repo: Updated user not found: %w", err)
	}

	return true, r.userCaching(ctx, &updatedUser)
}

func (r *UserRepository) PatchUser(ctx context.Context, req *models.PatchUserRequest) (bool, error) {
	result := r.db.Where("id = ?", req.Login).Model(&models.User{}).Updates(req)

	if result.Error != nil {
		return false, fmt.Errorf("Repo: User not found: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return false, fmt.Errorf("Repo: User not found: %w", result.Error)
	}

	var patchedUser models.User
	if err := r.db.Where("id = ?", req.Login).First(&patchedUser).Error; err != nil {
		return false, fmt.Errorf("Repo: Patched user not found: %w", err)
	}

	return true, r.userCaching(ctx, &patchedUser)
}


func (r *UserRepository) DeleteUser(ctx context.Context, req *models.DeleteUserRequest) (bool, error) {
	if err := r.db.Delete(&models.User{}, req.Login).Error; err != nil {
		return false, fmt.Errorf("Repo: User not found: %w", err)
	}

	
	r.redisClient.Del(ctx, req.Login).Err()

	return true, nil
}


