package handler

import (
	"net/http"

	"task-api/internal/usecase"
)

type TaskHandler interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	usecase usecase.TaskUsecase
}

func NewTaskHandler(usecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{usecase: usecase}
}

func (h *taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a task
}
