// package handler_test

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/interfaces/http/handler"
// 	"github.com/dimasafifudin11/golang-ddd-pattern/internal/domain/model"
// 	"github.com/gofiber/fiber/v3"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // MockAuthService is a mock of AuthService
// type MockAuthService struct {
// 	mock.Mock
// }

// func (m *MockAuthService) Register(ctx context.Context, name, email, password string) (*model.User, error) {
// 	args := m.Called(ctx, name, email, password)
// 	if args.Get(0) == nil {
// 		return nil, args.Error(1)
// 	}
// 	return args.Get(0).(*model.User), args.Error(1)
// }

// func (m *MockAuthService) Login(ctx context.Context, email, password string) (string, error) {
// 	args := m.Called(ctx, email, password)
// 	return args.String(0), args.Error(1)
// }

// func TestAuthHandler_Register(t *testing.T) {
// 	// Setup
// 	mockService := new(MockAuthService)
// 	authHandler := handler.NewAuthHandler(mockService)
// 	app := fiber.New()
// 	app.Post("/register", authHandler.Register)

// 	// Mock expectation
// 	// NOTE: In a real test, you'd create an actual user model instance.
// 	mockService.On("Register", mock.Anything, "Test User", "test@example.com", "password123").Return(&model.User{Name: "Test User"}, nil)

// 	// Create request
// 	requestBody := handler.RegisterRequest{
// 		Name:     "Test User",
// 		Email:    "test@example.com",
// 		Password: "password123",
// 	}
// 	bodyBytes, _ := json.Marshal(requestBody)
// 	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(bodyBytes))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Execute request
// 	resp, err := app.Test(req)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusCreated, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/interfaces/http/handler"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthService adalah mock dari AuthService
type MockAuthService struct {
	mock.Mock
}

// Implementasi method Register dari interface service.AuthService
func (m *MockAuthService) Register(ctx context.Context, name, email, password string) (*model.User, error) {
	args := m.Called(ctx, name, email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

// Implementasi method Login dari interface service.AuthService
func (m *MockAuthService) Login(ctx context.Context, email, password string) (string, error) {
	args := m.Called(ctx, email, password)
	return args.String(0), args.Error(1)
}

func TestAuthHandler_Register(t *testing.T) {
	// Setup
	mockService := new(MockAuthService)
	// Pastikan kita menggunakan interface, bukan implementasi langsung
	var authSvc service.AuthService = mockService
	authHandler := handler.NewAuthHandler(authSvc)
	app := fiber.New()
	app.Post("/api/register", authHandler.Register)

	// Mock expectation
	expectedUser := &model.User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}
	mockService.On("Register", mock.Anything, "Test User", "test@example.com", "password123").Return(expectedUser, nil)

	// Create request
	requestBody := handler.RegisterRequest{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	// Verifikasi bahwa method mock dipanggil
	mockService.AssertExpectations(t)
}
