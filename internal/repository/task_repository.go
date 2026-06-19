package repository

import (
	"context"
	"task-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Task, error)
	FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*domain.Task, error)
	FindAll(ctx context.Context) ([]*domain.Task, error)
	FindByStatus(ctx context.Context, status string) ([]*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id primitive.ObjectID) error
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

func (r *taskRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Task, error) {
	var task domain.Task
	err := r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*domain.Task, error) {
	cursor, err := r.collection.Find(ctx, primitive.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*domain.Task
	for cursor.Next(ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) FindAll(ctx context.Context) ([]*domain.Task, error) {
	cursor, err := r.collection.Find(ctx, primitive.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*domain.Task
	for cursor.Next(ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) FindByStatus(ctx context.Context, status string) ([]*domain.Task, error) {
	cursor, err := r.collection.Find(ctx, primitive.M{"status": status})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*domain.Task
	for cursor.Next(ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) Update(ctx context.Context, task *domain.Task) error {
	_, err := r.collection.UpdateOne(ctx, primitive.M{"_id": task.ID}, primitive.M{"$set": task})
	return err
}

func (r *taskRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, primitive.M{"_id": id})
	return err
}
