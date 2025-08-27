package handler

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/common"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetCurrent(c fiber.Ctx) error {
	// Ambil user_id dari context yang sudah di-set oleh middleware Auth
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(fiber.StatusUnauthorized, "Unauthorized"))
	}

	user, err := h.userService.GetProfile(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, user))
}
