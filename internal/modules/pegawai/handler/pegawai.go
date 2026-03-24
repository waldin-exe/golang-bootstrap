package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/entity"
	appErrors "github.com/waldin-exe/golang-bootstrap/utils/errors"
	"github.com/waldin-exe/golang-bootstrap/utils/response"
)

type PegawaiHandler struct {
	service contract.PegawaiService
}

func NewPegawaiHandler(service contract.PegawaiService) *PegawaiHandler {
	return &PegawaiHandler{service: service}
}

func (h *PegawaiHandler) GetPegawais(c *fiber.Ctx) error {
	var req entity.GetPegawaiRequest

	req.Limit = c.QueryInt("limit", 10)
	req.Offset = c.QueryInt("offset", 0)
	req.Id = c.QueryInt("id", 0)
	req.NamaPegawai = c.Query("nama_pegawai", "")
	req.JenisPegawai = c.Query("jenis_pegawai", "")
	req.TanggalBerangkat = c.Query("tanggal_berangkat", "")
	req.TanggalKembali = c.Query("tanggal_kembali", "")

	excludeIDsParam := c.Query("exclude_id", "")
	if excludeIDsParam != "" {
		idStrs := strings.Split(excludeIDsParam, ",")
		for _, idStr := range idStrs {
			if id, err := strconv.Atoi(strings.TrimSpace(idStr)); err == nil {
				req.ExcludeIDs = append(req.ExcludeIDs, id)
			}
		}
	}

	data, totalData, err := h.service.GetPegawais(c.Context(), req)
	if err != nil {
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	payload := map[string]interface{}{
		"rows":       data,
		"total_data": totalData,
	}

	return response.NewResponseSuccess(c, payload, "Success")
}

func (h *PegawaiHandler) CreatePegawai(c *fiber.Ctx) error {
	var req entity.PostPegawaiRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	userID := fmt.Sprintf("%v", c.Locals("user_id"))
	req.CreatedBy = userID

	if err := h.service.CreatePegawai(c.Context(), req); err != nil {
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseCreated(c, nil, "Pegawai & user berhasil ditambahkan")
}

func (h *PegawaiHandler) UpdatePegawai(c *fiber.Ctx) error {
	var req entity.PutPegawaiRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, "invalid request body")
	}

	userID := fmt.Sprintf("%v", c.Locals("user_id"))
	req.UpdatedBy = userID

	if req.Id == 0 {
		return response.NewResponseBadRequest(c, "id is required")
	}

	if err := h.service.UpdatePegawai(c.Context(), req); err != nil {
		if appErrors.IsNotFoundError(err) {
			return response.NewResponseError(c, fiber.StatusNotFound, err.Error())
		}
		if appErrors.IsBadRequestError(err) {
			return response.NewResponseBadRequest(c, err.Error())
		}
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseSuccess(c, nil, "Pegawai berhasil diperbarui")
}

func (h *PegawaiHandler) DeletePegawai(c *fiber.Ctx) error {
	var req entity.DeletePegawaiRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, "invalid request body")
	}

	userID := fmt.Sprintf("%v", c.Locals("user_id"))
	req.UpdatedBy = userID

	if err := h.service.DeletePegawai(c.Context(), req); err != nil {
		if appErrors.IsNotFoundError(err) {
			return response.NewResponseError(c, fiber.StatusNotFound, err.Error())
		}
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseSuccess(c, nil, "Pegawai berhasil dihapus")
}
