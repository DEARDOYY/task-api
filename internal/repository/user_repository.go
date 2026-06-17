package repository

import (
	"context"
	"task-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	// เอา id ที่ MongoDB สร้างให้ ใส่กลับเข้า struct
	user.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}
