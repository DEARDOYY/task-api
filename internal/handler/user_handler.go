package handler

import (
	"task-api/internal/usecase"
	"task-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.usecase.GetUserByID(c.Request.Context(), id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.OK(c, user)
}
