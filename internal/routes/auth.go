package routes

import (
	"task-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoute(rg *gin.RouterGroup) {
	user := rg.Group("/auth")
	{
		user.POST("/register", handler.NewAuthHandler(nil).Register)
	}
}
