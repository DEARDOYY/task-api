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
	GetTaskByID(ctx context.Context, id string) (*domain.Task, error)
	GetTasksByUserID(ctx context.Context, userID string) ([]*domain.Task, error)
	GetAllTasks(ctx context.Context) ([]*domain.Task, error)
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

func (u *taskUsecase) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task id")
	}

	task, err := u.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (u *taskUsecase) GetTasksByUserID(ctx context.Context, userID string) ([]*domain.Task, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	tasks, err := u.repo.FindByUserID(ctx, userObjID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (u *taskUsecase) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (u *taskUsecase) GetTasksByStatus(ctx context.Context, status string) ([]*domain.Task, error) {
	tasks, err := u.repo.FindByStatus(ctx, status)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (u *taskUsecase) UpdateTaskStatus(ctx context.Context, id string, status string) (*domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task id")
	}

	task, err := u.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	task.Status = status
	task.UpdatedAt = time.Now()

	if err := u.repo.Update(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (u *taskUsecase) DeleteTask(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task id")
	}

	if err := u.repo.Delete(ctx, objID); err != nil {
		return err
	}
	return nil
}
