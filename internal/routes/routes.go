package routes

import (
	"task-api/internal"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *internal.Handlers) {
	api := r.Group("/api/v1")

	TaskRoute(api)
	UserRoute(api, h.User)
	AuthRoute(api, h.Auth)

}
