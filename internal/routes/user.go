package routes

import (
	"task-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRoute(rg *gin.RouterGroup) {
	user := rg.Group("/users")
	{
		user.POST("/register", handler.NewUserHandler(nil).Register)
	}
}
