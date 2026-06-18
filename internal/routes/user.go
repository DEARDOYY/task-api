package routes

import (
	"task-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRoute(rg *gin.RouterGroup, h *handler.UserHandler) {
	user := rg.Group("/users")
	{
		user.GET("/:id", h.GetUser)
	}
}
