package repository

import (
	"context"
	"task-api/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository { // 👈 return เป็น interface
	return &userRepository{collection: collection}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}
