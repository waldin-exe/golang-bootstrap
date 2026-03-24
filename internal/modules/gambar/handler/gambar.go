package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/entity"
	appErrors "github.com/waldin-exe/golang-bootstrap/utils/errors"
	"github.com/waldin-exe/golang-bootstrap/utils/response"
)

type GambarHandler struct {
	service contract.GambarService
}

func NewGambarHandler(service contract.GambarService) *GambarHandler {
	return &GambarHandler{service: service}
}

func (h *GambarHandler) GetGambar(c *fiber.Ctx) error {
	var req entity.GetGambarRequest
	req.ForeignID = c.QueryInt("foreign_id", 0)
	req.Reference = c.Query("reference", "")

	data, err := h.service.GetGambarByReference(c.Context(), req)
	if err != nil {
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	payload := map[string]interface{}{
		"rows": data,
	}

	return response.NewResponseSuccess(c, payload, "Success")
}

func (h *GambarHandler) UploadGambar(c *fiber.Ctx) error {
	foreignID := c.FormValue("foreign_id")
	reference := c.FormValue("reference")

	if foreignID == "" || reference == "" {
		return response.NewResponseBadRequest(c, "foreign_id dan reference wajib diisi")
	}

	fIDInt := 0
	if _, err := fmt.Sscanf(foreignID, "%d", &fIDInt); err != nil {
		return response.NewResponseBadRequest(c, "foreign_id harus berupa angka")
	}

	form, err := c.MultipartForm()
	if err != nil || form == nil {
		return response.NewResponseBadRequest(c, "Form upload tidak valid atau tidak ditemukan")
	}

	files := form.File["images"]
	if len(files) == 0 {
		return response.NewResponseBadRequest(c, "Tidak ada file yang diunggah")
	}

	// Call service
	if err := h.service.UploadImages(c.Context(), reference, fIDInt, files); err != nil {
		if appErrors.IsBadRequestError(err) {
			return response.NewResponseBadRequest(c, err.Error())
		}
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseCreated(c, nil, "Gambar berhasil ditambahkan")
}

func (h *GambarHandler) DeleteGambar(c *fiber.Ctx) error {
	var req entity.DeleteGambarRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, "Format JSON tidak valid")
	}

	if len(req.DeleteImageIds) == 0 {
		return response.NewResponseBadRequest(c, "delete_image_ids wajib diisi")
	}

	if err := h.service.DeleteGambar(c.Context(), req); err != nil {
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseSuccess(c, nil, "Gambar berhasil dihapus")
}
