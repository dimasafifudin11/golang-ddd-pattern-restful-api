package repository

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
)

type ContactRepository interface {
	Save(ctx context.Context, contact *model.Contact) error
	FindByID(ctx context.Context, contactID uint) (*model.Contact, error)
	FindAllByUserID(ctx context.Context, userID uint) ([]model.Contact, error)
	Update(ctx context.Context, contact *model.Contact) error
	Delete(ctx context.Context, contactID uint) error
}
