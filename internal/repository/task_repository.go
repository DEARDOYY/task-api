package repository

import (
	"context"
	"task-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
}

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) TaskRepository { // 👈 return เป็น interface
	return &taskRepository{collection: collection}
}

func (r *taskRepository) Create(ctx context.Context, task *domain.Task) error {
	result, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		return err
	}

	task.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}
