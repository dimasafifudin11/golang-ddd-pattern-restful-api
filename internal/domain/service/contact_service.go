package service

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
)

type ContactCreateRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone"`
}

type ContactUpdateRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone"`
}

type ContactService interface {
	CreateContact(ctx context.Context, userID uint, request ContactCreateRequest) (*model.Contact, error)
	GetContactByID(ctx context.Context, contactID uint) (*model.Contact, error)
	GetAllContactsByUserID(ctx context.Context, userID uint) ([]model.Contact, error)
	UpdateContact(ctx context.Context, contactID uint, request ContactUpdateRequest) (*model.Contact, error)
	DeleteContact(ctx context.Context, contactID uint) error
}
