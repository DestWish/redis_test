package main

import (
	"context"
	"fmt"

	"github.com/DestWish/redis_test/internal/handler"
	"github.com/DestWish/redis_test/internal/models"
	"github.com/DestWish/redis_test/internal/repository"
	"github.com/DestWish/redis_test/internal/service"
	"github.com/redis/go-redis/v9"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := initDB()
	if err != nil {
		fmt.Printf("Failed to init database: %v", err)
	}

	cache := initCache()
	defer cache.Close()

	ctx := context.Background()

	userRepo := repository.New_userRepo(db, cache)

	userService := service.New_userService(userRepo, cache)

	userHandler := handler.New_userHandler(userService)

	userId := userHandler.Create(ctx, models.Create_userRequest{Email: "testMail", Name: "Obezjana"})

	fmt.Printf("был создан Юзер, ID: %v \n", userId)

	fmt.Printf("Получим из кэша юзера %v \n", userHandler.Get(ctx, userId))
}


func initCache() (*redis.Client) {
	rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
    Protocol: 2,
  })
  
//docker run -d --name redis -p 6379:6379 redis:8.6.2


  return rdb
}


func initDB() (*gorm.DB, error) {
	//docker run -d --name cards -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=cards -p 5432:5432 postgres:16-alpine
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=cards sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed connect to db: %w", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("failed to migrate %w", err)
	}

	return db, nil
}