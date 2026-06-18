package usecase

import (
	"context"
	"errors"

	"task-api/internal/domain"
	"task-api/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type UserUsecase interface {
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*domain.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	user, err := u.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return u.repo.FindAll(ctx)
}

func (u *userUsecase) UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	err = u.repo.Update(ctx, objID, user)
	if err != nil {
		return nil, errors.New("failed to update user")
	}

	return u.repo.FindByID(ctx, objID)
}
