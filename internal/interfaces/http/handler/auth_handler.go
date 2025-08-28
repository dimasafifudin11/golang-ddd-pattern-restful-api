package handler

import (
	"fmt"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/common"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/pkg/util"
	"github.com/gofiber/fiber/v3"
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

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Request body tidak valid"))
	}

	if validationErrors := util.ValidateStruct(req); validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "Bad Request",
			"errors": validationErrors,
		})
	}

	user, err := h.authService.Register(c, req.Name, req.Email, req.Password)
	if err != nil {
		switch err {
		case common.ErrConflict:
			return c.Status(fiber.StatusConflict).JSON(common.NewErrorResponse(fiber.StatusConflict, "Email sudah terdaftar"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal melakukan registrasi"))
		}
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(fiber.StatusCreated, user))
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Request body tidak valid"))
	}

	if validationErrors := util.ValidateStruct(req); validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "Bad Request",
			"errors": validationErrors,
		})
	}

	token, err := h.authService.Login(c, req.Email, req.Password)
	if err != nil {
		switch err {
		case common.ErrUnauthorized:
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(fiber.StatusUnauthorized, "Email atau password salah"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal melakukan login"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, fiber.Map{"token": token}))
}
