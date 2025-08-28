package handler

import (
	"strconv"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/common"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/pkg/util"
	"github.com/gofiber/fiber/v3"
)

type AddressHandler struct {
	addressService service.AddressService
}

func NewAddressHandler(addressService service.AddressService) *AddressHandler {
	return &AddressHandler{addressService: addressService}
}

func (h *AddressHandler) Create(c fiber.Ctx) error {
	contactID, err := strconv.Atoi(c.Params("contactId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID Kontak tidak valid"))
	}

	var req service.AddressCreateRequest
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

	address, err := h.addressService.CreateAddress(c, uint(contactID), req)
	if err != nil {
		switch err {
		case common.ErrNotFound:
			return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, "Kontak pemilik alamat tidak ditemukan"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal membuat alamat"))
		}
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(fiber.StatusCreated, address))
}

func (h *AddressHandler) GetByID(c fiber.Ctx) error {
	addressID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID Alamat tidak valid"))
	}

	address, err := h.addressService.GetAddressByID(c, uint(addressID))
	if err != nil {
		switch err {
		case common.ErrNotFound:
			return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, "Alamat dengan ID tersebut tidak ditemukan"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal mengambil data alamat"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, address))
}

func (h *AddressHandler) GetAllByContactID(c fiber.Ctx) error {
	contactID, err := strconv.Atoi(c.Params("contactId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID Kontak tidak valid"))
	}

	addresses, err := h.addressService.GetAllAddressesByContactID(c, uint(contactID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal mengambil daftar alamat"))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, addresses))
}

func (h *AddressHandler) Update(c fiber.Ctx) error {
	addressID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID Alamat tidak valid"))
	}

	var req service.AddressUpdateRequest
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

	address, err := h.addressService.UpdateAddress(c, uint(addressID), req)
	if err != nil {
		switch err {
		case common.ErrNotFound:
			return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, "Alamat dengan ID tersebut tidak ditemukan untuk diupdate"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal mengupdate alamat"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, address))
}

func (h *AddressHandler) Delete(c fiber.Ctx) error {
	addressID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "ID Alamat tidak valid"))
	}

	if err := h.addressService.DeleteAddress(c, uint(addressID)); err != nil {
		switch err {
		case common.ErrNotFound:
			return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(fiber.StatusNotFound, "Alamat dengan ID tersebut tidak ditemukan untuk dihapus"))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError, "Gagal menghapus alamat"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(fiber.StatusOK, fiber.Map{"message": "Alamat berhasil dihapus"}))
}
