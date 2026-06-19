package internal

import (
	"os"

	"task-api/internal/handler"
	"task-api/internal/repository"
	"task-api/internal/usecase"
	"task-api/pkg/database"
)

type Handlers struct {
	Auth *handler.AuthHandler
	User *handler.UserHandler
	Task *handler.TaskHandler
}

func InitHandlers() *Handlers {
	userCollection := database.GetCollection(os.Getenv("MONGO_DB_NAME"), "users")
	taskCollection := database.GetCollection(os.Getenv("MONGO_DB_NAME"), "tasks")

	userRepo := repository.NewUserRepository(userCollection)
	taskRepo := repository.NewTaskRepository(taskCollection)

	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	return &Handlers{
		Auth: authHandler,
		User: userHandler,
		Task: taskHandler,
	}
}
