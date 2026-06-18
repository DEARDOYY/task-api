package routes

import (
	"task-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRoute(rg *gin.RouterGroup, h *handler.UserHandler) {
	user := rg.Group("/users")
	{
		// เพิ่ม route อื่นๆ เช่น POST /users, PUT /users/:id, DELETE /users/:id ตามต้องการ
		user.GET("/:id", h.GetUser)
		user.GET("/", h.GetAllUsers)
		user.PUT("/:id", h.UpdateUser)

	}
}
