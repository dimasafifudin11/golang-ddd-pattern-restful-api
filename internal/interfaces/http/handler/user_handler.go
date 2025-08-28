package handler

import (
	"strconv"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/common"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/pkg/util"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetCurrent(c fiber.Ctx) error {
	// ... (implementation is the same as before)
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(fiber.StatusUnauthorized, "Unauthorized"))
	}
	user, err := h.userService.GetProfile(c, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, user))
}

func (h *UserHandler) GetAll(c fiber.Ctx) error {
	// ... (implementation is the same as before)
	users, err := h.userService.GetAllUsers(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, users))
}

func (h *UserHandler) GetByID(c fiber.Ctx) error {
	// ... (implementation is the same as before)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid ID format"))
	}
	user, err := h.userService.GetUserByID(c, uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, user))
}

func (h *UserHandler) Update(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID user tidak valid"))
	}

	var req service.UserUpdateRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Request body tidak valid"))
	}
	if validationErrors := util.ValidateStruct(req); validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.StatusBadRequest, "status": "Bad Request", "errors": validationErrors})
	}

	user, err := h.userService.UpdateUser(c, uint(id), req)
	if err != nil {
		switch err {
		case common.ErrNotFound:
			return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, "User dengan ID tersebut tidak ditemukan"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal mengupdate data user"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, user))
}

func (h *UserHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID user tidak valid"))
	}

	err = h.userService.DeleteUser(c, uint(id))
	if err != nil {
		switch err {
		case common.ErrNotFound:
			return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, "User dengan ID tersebut tidak ditemukan"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal menghapus user"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, fiber.Map{"message": "User berhasil dihapus"}))
}
