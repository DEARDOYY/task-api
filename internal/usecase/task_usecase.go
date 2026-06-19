package usecase

import (
	"context"

	"task-api/internal/domain"
	"task-api/internal/repository"
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
	task := &domain.Task{
		Title:  req.Title,
		Status: req.Status,
		UserID: req.UserID,
	}

	err := u.repo.Create(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
