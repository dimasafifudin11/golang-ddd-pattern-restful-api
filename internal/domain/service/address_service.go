package service

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
)

type AddressCreateRequest struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country" validate:"required"`
	PostalCode string `json:"postal_code"`
}

type AddressUpdateRequest struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country" validate:"required"`
	PostalCode string `json:"postal_code"`
}

type AddressService interface {
	CreateAddress(ctx context.Context, contactID uint, request AddressCreateRequest) (*model.Address, error)
	GetAddressByID(ctx context.Context, addressID uint) (*model.Address, error)
	GetAllAddressesByContactID(ctx context.Context, contactID uint) ([]model.Address, error)
	UpdateAddress(ctx context.Context, addressID uint, request AddressUpdateRequest) (*model.Address, error)
	DeleteAddress(ctx context.Context, addressID uint) error
}
