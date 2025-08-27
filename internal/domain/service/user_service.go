package service

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
)

// UserService defines the interface for user profile related logic.
type UserService interface {
	GetProfile(ctx context.Context, userID uint) (*model.User, error)
}
