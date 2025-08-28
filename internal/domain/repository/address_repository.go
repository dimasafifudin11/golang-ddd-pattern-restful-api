package repository

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
)

type AddressRepository interface {
	Save(ctx context.Context, address *model.Address) error
	FindByID(ctx context.Context, addressID uint) (*model.Address, error)
	FindAllByContactID(ctx context.Context, contactID uint) ([]model.Address, error)
	Update(ctx context.Context, address *model.Address) error
	Delete(ctx context.Context, addressID uint) error
}
