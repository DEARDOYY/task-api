package usecase

import (
	"context"
	"errors"

	"task-api/internal/domain"
	"task-api/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
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
