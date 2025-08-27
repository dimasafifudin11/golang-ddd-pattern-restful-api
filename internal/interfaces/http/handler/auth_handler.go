package handler

import (
	"fmt"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/common"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/pkg/util"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/net/context"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid request body"))
	}

	if err := util.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	user, err := h.authService.Register(context.Background(), req.Name, req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(fiber.StatusCreated, user))
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid request body"))
	}

	if err := util.ValidateStruct(req); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	token, err := h.authService.Login(context.Background(), req.Email, req.Password)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(fiber.StatusUnauthorized, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, fiber.Map{"token": token}))
}
