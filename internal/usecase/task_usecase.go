package usecase

import (
	"context"
	"errors"
	"time"

	"task-api/internal/domain"
	"task-api/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTaskRequest struct {
	Title  string `json:"title" binding:"required"`
	Status string `json:"status"`
	UserID string `json:"user_id" binding:"required"`
}

type TaskUsecase interface {
	CreateTask(ctx context.Context, req CreateTaskRequest) (*domain.Task, error)
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: repo}
}

func (u *taskUsecase) CreateTask(ctx context.Context, req CreateTaskRequest) (*domain.Task, error) {
	userObjID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	task := &domain.Task{
		Title:     req.Title,
		Status:    "pending",
		UserID:    userObjID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.repo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}
