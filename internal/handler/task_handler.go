package handler

import (
	"task-api/internal/usecase"
	"task-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	usecase usecase.TaskUsecase
}

func NewTaskHandler(usecase usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{usecase: usecase}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req usecase.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	task, err := h.usecase.CreateTask(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Created(c, task)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.usecase.GetTaskByID(c.Request.Context(), id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, task)
}

func (h *TaskHandler) GetTasksByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	tasks, err := h.usecase.GetTasksByUserID(c.Request.Context(), userID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, tasks)
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.usecase.GetAllTasks(c.Request.Context())
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, tasks)
}
