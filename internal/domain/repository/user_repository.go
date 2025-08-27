package repository

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
)

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	Save(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByID(ctx context.Context, id uint) (*model.User, error)
}
