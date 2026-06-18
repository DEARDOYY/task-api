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
}

func InitHandlers() *Handlers {
	userCollection := database.GetCollection(os.Getenv("MONGO_DB_NAME"), "users")
	userRepo := repository.NewUserRepository(userCollection)
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// taskCollection := database.GetCollection(os.Getenv("MONGO_DB_NAME"), "tasks")
	// taskRepo := repository.NewTaskRepository(taskCollection)
	// taskUsecase := usecase.NewTaskUsecase(taskRepo)
	// taskHandler := handler.NewTaskHandler(taskUsecase)

	return &Handlers{
		Auth: authHandler,
		User: userHandler,
		// Task: taskHandler,
	}
}
