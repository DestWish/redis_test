package repository

import (
	"fmt"

	"github.com/DestWish/redis_test/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New_userRepo (db *gorm.DB) *UserRepository {
 return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	result := r.db.Create(user)
	if err := result.Error; err != nil {
		fmt.Printf("Error %v", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUser(userID uint) (models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", userID).Find(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}