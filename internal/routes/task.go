package routes

import (
	"task-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func TaskRoute(rg *gin.RouterGroup, h *handler.TaskHandler) {
	task := rg.Group("/tasks")
	{
		task.POST("/", h.CreateTask)
	}
}
