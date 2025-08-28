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

type MockAddressService struct {
	mock.Mock
}

func (m *MockAddressService) CreateAddress(ctx context.Context, contactID uint, request service.AddressCreateRequest) (*model.Address, error) {
	args := m.Called(ctx, contactID, request)
	return args.Get(0).(*model.Address), args.Error(1)
}
func (m *MockAddressService) GetAddressByID(ctx context.Context, addressID uint) (*model.Address, error) {
	args := m.Called(ctx, addressID)
	return args.Get(0).(*model.Address), args.Error(1)
}
func (m *MockAddressService) GetAllAddressesByContactID(ctx context.Context, contactID uint) ([]model.Address, error) {
	args := m.Called(ctx, contactID)
	return args.Get(0).([]model.Address), args.Error(1)
}
func (m *MockAddressService) UpdateAddress(ctx context.Context, addressID uint, request service.AddressUpdateRequest) (*model.Address, error) {
	args := m.Called(ctx, addressID, request)
	return args.Get(0).(*model.Address), args.Error(1)
}
func (m *MockAddressService) DeleteAddress(ctx context.Context, addressID uint) error {
	args := m.Called(ctx, addressID)
	return args.Error(0)
}

func TestAddressHandler_GetByID(t *testing.T) {
	mockService := new(MockAddressService)
	addressHandler := handler.NewAddressHandler(mockService)

	app := fiber.New()
	app.Get("/api/addresses/:id", addressHandler.GetByID)

	expectedAddress := &model.Address{ID: 1, Country: "Indonesia"}
	mockService.On("GetAddressByID", mock.Anything, uint(1)).Return(expectedAddress, nil)

	req, _ := http.NewRequest("GET", "/api/addresses/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}
