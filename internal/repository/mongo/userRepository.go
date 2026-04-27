package mongo

import (
	"context"
	"fmt"

	"github.com/DestWish/redis_test/internal/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	mongoClient *mongo.Client
	redisClient *redis.Client
}

func NewUserRepo(mongoClient *mongo.Client, redisClient *redis.Client) *UserRepository {
	return &UserRepository{mongoClient: mongoClient, redisClient: redisClient}
}

func (r *UserRepository) userCaching(ctx context.Context, user *models.User) error {
	key := userCacheKey(user.Login)
	if err := r.redisClient.HSet(ctx, key, user).Err(); err != nil {
		return fmt.Errorf("Repo: Cache failed: %w", err)
	}

	return nil
}

func (r *UserRepository) CreateUser(ctx context.Context, req *models.CreateUserRequest) (string, error) {
	user := &models.User{Login: req.Login, Email: req.Email, Name: req.Name}
	_, err := r.mongoClient.Database("mongotest").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("Repo: Create user failed: %w", err)
	}

	return user.Login, r.userCaching(ctx, user)
}

func (r *UserRepository) GetUser(ctx context.Context, req *models.ReadUserRequest) (models.User, error) {
	key := userCacheKey(req.Login)
	var user models.User

	err := r.redisClient.HGetAll(ctx, key).Scan(&user)
	if err == nil && user.Login != "" {
		return user, nil
	}

	if  err := r.mongoClient.Database("mongotest").Collection("users").FindOne(ctx, bson.M{"_id": req.Login}).Decode(&user); err != nil {
		return user, fmt.Errorf("Repo: User not found: %w", err)
	}

	return user, r.userCaching(ctx, &user)
}

func (r *UserRepository) ReplaceUser(ctx context.Context, req *models.UpdateUserRequest) (bool, error) {
	if err := r.mongoClient.Database("mongotest").Collection("users").FindOneAndReplace(ctx, bson.M{"_id": req.Login}, req).Err(); err != nil {
		return false, fmt.Errorf("Repo: User not found: %w", err)
	}

	var updatedUser models.User
	if err := r.mongoClient.Database("mongotest").Collection("users").FindOne(ctx, bson.M{"_id": req.Login}).Decode(&updatedUser); err != nil {
		return false, fmt.Errorf("Repo: Updated user not found: %w", err)
	}

	return true, r.userCaching(ctx, &updatedUser)
}

func (r *UserRepository) PatchUser(ctx context.Context, req *models.PatchUserRequest) (bool, error) {
	updateFields := bson.M{}
	if req.Email != nil {
		updateFields["email"] = *req.Email
	}
	if req.Name != nil {
		updateFields["name"] = *req.Name
	}
	
	if err := r.mongoClient.Database("mongotest").Collection("users").FindOneAndUpdate(ctx, bson.M{"_id": req.Login}, bson.M{"$set": updateFields}).Err(); err != nil {
		return false, fmt.Errorf("Repo: User not found: %w", err)
	}

	var patchedUser models.User
	if err := r.mongoClient.Database("mongotest").Collection("users").FindOne(ctx, bson.M{"_id": req.Login}).Decode(&patchedUser); err != nil {
		return false, fmt.Errorf("Repo: Updated user not found: %w", err)
	}

	return true, r.userCaching(ctx, &patchedUser)
}

func (r *UserRepository) DeleteUser(ctx context.Context, req *models.DeleteUserRequest) (bool, error) {
	if err := r.mongoClient.Database("mongotest").Collection("users").FindOneAndDelete(ctx, bson.M{"_id": req.Login}).Err(); err != nil {
		return false, fmt.Errorf("Repo: User not found: %w", err)
	}

	key := userCacheKey(req.Login)
	r.redisClient.Del(ctx, key).Err()

	return true, nil
}

func userCacheKey(userLogin string) string {
	return fmt.Sprintf("user:%v", userLogin)
}
