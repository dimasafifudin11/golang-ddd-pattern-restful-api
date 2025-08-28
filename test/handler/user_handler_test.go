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

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetProfile(ctx context.Context, userID uint) (*model.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*model.User), args.Error(1)
}
func (m *MockUserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.User), args.Error(1)
}
func (m *MockUserService) GetUserByID(ctx context.Context, userID uint) (*model.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*model.User), args.Error(1)
}
func (m *MockUserService) UpdateUser(ctx context.Context, userID uint, request service.UserUpdateRequest) (*model.User, error) {
	args := m.Called(ctx, userID, request)
	return args.Get(0).(*model.User), args.Error(1)
}
func (m *MockUserService) DeleteUser(ctx context.Context, userID uint) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func TestUserHandler_GetAll(t *testing.T) {
	mockService := new(MockUserService)
	userHandler := handler.NewUserHandler(mockService)

	app := fiber.New()
	app.Get("/api/users", userHandler.GetAll)

	expectedUsers := []model.User{{ID: 1, Name: "Test User"}}
	mockService.On("GetAllUsers", mock.Anything).Return(expectedUsers, nil)

	req, _ := http.NewRequest("GET", "/api/users", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}
