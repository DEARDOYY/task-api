package routes

import (
	"task-api/internal/handler"
	"task-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoute(rg *gin.RouterGroup, h *handler.TaskHandler) {
	task := rg.Group("/tasks")
	task.Use(middleware.AuthRequired())
	{
		task.POST("/", h.CreateTask)
		task.GET("/:id", h.GetTaskByID)
		task.GET("/user/:user_id", h.GetTasksByUserID)
		task.GET("/", h.GetAllTasks)
		task.GET("/status/:status", h.GetTasksByStatus)
		task.PUT("/:id/status", h.UpdateTaskStatus)
		task.DELETE("/:id", h.DeleteTask)
	}
}
