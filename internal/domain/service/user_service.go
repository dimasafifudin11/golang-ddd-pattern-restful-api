package service

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
)

type UserUpdateRequest struct {
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
}

type UserService interface {
	GetProfile(ctx context.Context, userID uint) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, userID uint) (*model.User, error)
	UpdateUser(ctx context.Context, userID uint, request UserUpdateRequest) (*model.User, error)
	DeleteUser(ctx context.Context, userID uint) error
}
