package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/armada/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/armada/entity"
	appErrors "github.com/waldin-exe/golang-bootstrap/utils/errors"
	"github.com/waldin-exe/golang-bootstrap/utils/response"
	"mime/multipart"
)

type ArmadaHandler struct {
	service contract.ArmadaService
}

func NewArmadaHandler(service contract.ArmadaService) *ArmadaHandler {
	return &ArmadaHandler{service: service}
}

func (h *ArmadaHandler) GetArmadas(c *fiber.Ctx) error {
	var req entity.GetArmadaRequest
	req.Limit = c.QueryInt("limit", 10)
	req.Offset = c.QueryInt("offset", 0)
	req.Id = c.QueryInt("id", 0)
	req.PlatNomor = c.Query("plat_nomor", "")
	req.Merk = c.Query("merk", "")
	req.Jenis = c.Query("jenis", "")
	req.JumlahSeat = c.QueryInt("jumlah_seat", 0)

	data, totalData, err := h.service.GetArmadas(c.Context(), req)
	if err != nil {
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	payload := map[string]interface{}{
		"rows":       data,
		"total_data": totalData,
	}

	return response.NewResponseSuccess(c, payload, "Success")
}

func (h *ArmadaHandler) CreateArmada(c *fiber.Ctx) error {
	var req entity.PostArmadaRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	userID := fmt.Sprintf("%v", c.Locals("user_id"))
	req.CreatedBy = userID

	form, _ := c.MultipartForm()

	var files []*multipart.FileHeader
	if form != nil {
		files = form.File["images"]
	}

	if err := h.service.CreateArmada(c.Context(), req, files); err != nil {
		if appErrors.IsBadRequestError(err) {
			return response.NewResponseBadRequest(c, err.Error())
		}
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseCreated(c, nil, "Armada berhasil ditambahkan")
}

func (h *ArmadaHandler) UpdateArmada(c *fiber.Ctx) error {
	var req entity.PutArmadaRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	userID := fmt.Sprintf("%v", c.Locals("user_id"))
	req.UpdatedBy = userID

	form, _ := c.MultipartForm()

	var files []*multipart.FileHeader
	if form != nil {
		files = form.File["images"]
	}
	var deletedImages []int
	if form != nil && len(form.Value["delete_image_ids"]) > 0 {
		str := form.Value["delete_image_ids"][0]
		_ = json.Unmarshal([]byte(str), &deletedImages)
	}

	if err := h.service.UpdateArmada(c.Context(), req, files, deletedImages); err != nil {
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseSuccess(c, nil, "Armada berhasil diperbarui")
}

func (h *ArmadaHandler) DeleteArmada(c *fiber.Ctx) error {
	var req entity.DeleteArmadaRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	userID := fmt.Sprintf("%v", c.Locals("user_id"))
	req.UpdatedBy = userID

	if err := h.service.DeleteArmada(c.Context(), req); err != nil {
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseSuccess(c, nil, "Armada berhasil dihapus")
}

func (h *ArmadaHandler) GetArmadasTersedia(c *fiber.Ctx) error {
	tglMulai := c.Query("tgl_mulai")
	tglAkhir := c.Query("tgl_akhir")

	if tglMulai == "" || tglAkhir == "" {
		return response.NewResponseBadRequest(c, "tgl_mulai dan tgl_akhir wajib diisi")
	}

	req := entity.GetArmadaTersediaRequest{
		TanggalMulai: tglMulai,
		TanggalAkhir: tglAkhir,
		Limit:        c.QueryInt("limit", 10),
		Offset:       c.QueryInt("offset", 0),
	}

	data, totalData, err := h.service.GetArmadasTersedia(c.Context(), req)
	if err != nil {
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	payload := map[string]interface{}{
		"rows":       data,
		"total_data": totalData,
	}

	return response.NewResponseSuccess(c, payload, "Success")
}
