package handler_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/interfaces/http/handler"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockContactService struct {
	mock.Mock
}

func (m *MockContactService) CreateContact(ctx context.Context, userID uint, request service.ContactCreateRequest) (*model.Contact, error) {
	args := m.Called(ctx, userID, request)
	return args.Get(0).(*model.Contact), args.Error(1)
}
func (m *MockContactService) GetContactByID(ctx context.Context, contactID uint) (*model.Contact, error) {
	args := m.Called(ctx, contactID)
	return args.Get(0).(*model.Contact), args.Error(1)
}
func (m *MockContactService) GetAllContactsByUserID(ctx context.Context, userID uint) ([]model.Contact, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]model.Contact), args.Error(1)
}
func (m *MockContactService) UpdateContact(ctx context.Context, contactID uint, request service.ContactUpdateRequest) (*model.Contact, error) {
	args := m.Called(ctx, contactID, request)
	return args.Get(0).(*model.Contact), args.Error(1)
}
func (m *MockContactService) DeleteContact(ctx context.Context, contactID uint) error {
	args := m.Called(ctx, contactID)
	return args.Error(0)
}

func TestContactHandler_GetByID(t *testing.T) {
	mockService := new(MockContactService)
	contactHandler := handler.NewContactHandler(mockService)

	app := fiber.New()
	app.Get("/api/contacts/:id", contactHandler.GetByID)

	expectedContact := &model.Contact{ID: 1, FirstName: "Test"}
	mockService.On("GetContactByID", mock.Anything, uint(1)).Return(expectedContact, nil)

	req, _ := http.NewRequest("GET", "/api/contacts/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}
