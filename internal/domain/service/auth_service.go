package service

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
)

// AuthService defines the interface for authentication and user management logic.
type AuthService interface {
	Register(ctx context.Context, name, email, password string) (*model.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}
