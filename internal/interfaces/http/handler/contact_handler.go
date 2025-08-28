package handler

import (
	"strconv"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/common"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/pkg/util"
	"github.com/gofiber/fiber/v3"
)

type ContactHandler struct {
	contactService service.ContactService
}

func NewContactHandler(contactService service.ContactService) *ContactHandler {
	return &ContactHandler{contactService: contactService}
}

func (h *ContactHandler) Create(c fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID User tidak valid"))
	}

	var req service.ContactCreateRequest
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

	contact, err := h.contactService.CreateContact(c, uint(userID), req)
	if err != nil {
		switch err {
		case common.ErrNotFound:
			return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, "User pemilik kontak tidak ditemukan"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal membuat kontak"))
		}
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(fiber.StatusCreated, contact))
}

func (h *ContactHandler) GetByID(c fiber.Ctx) error {
	contactID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID Kontak tidak valid"))
	}

	contact, err := h.contactService.GetContactByID(c, uint(contactID))
	if err != nil {
		switch err {
		case common.ErrNotFound:
			return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, "Kontak dengan ID tersebut tidak ditemukan"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal mengambil data kontak"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, contact))
}

func (h *ContactHandler) GetAllByUserID(c fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID User tidak valid"))
	}

	contacts, err := h.contactService.GetAllContactsByUserID(c, uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal mengambil daftar kontak"))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, contacts))
}

func (h *ContactHandler) Update(c fiber.Ctx) error {
	contactID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID Kontak tidak valid"))
	}

	var req service.ContactUpdateRequest
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

	contact, err := h.contactService.UpdateContact(c, uint(contactID), req)
	if err != nil {
		switch err {
		case common.ErrNotFound:
			return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, "Kontak dengan ID tersebut tidak ditemukan untuk diupdate"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal mengupdate kontak"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, contact))
}

func (h *ContactHandler) Delete(c fiber.Ctx) error {
	contactID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID Kontak tidak valid"))
	}

	if err := h.contactService.DeleteContact(c, uint(contactID)); err != nil {
		switch err {
		case common.ErrNotFound:
			return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, "Kontak dengan ID tersebut tidak ditemukan untuk dihapus"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal menghapus kontak"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, fiber.Map{"message": "Kontak berhasil dihapus"}))
}
