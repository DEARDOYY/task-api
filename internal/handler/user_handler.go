package handler

import (
	"task-api/internal/usecase"
	"task-api/pkg/response"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetUserByID(c *gin.Context, id string) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}

	user, err := h.usecase.GetUserByID(c.Request.Context(), objectID.Hex())
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, user)
}
