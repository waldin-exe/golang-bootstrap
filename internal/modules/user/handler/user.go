package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/user/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/user/entity"
	appErrors "github.com/waldin-exe/golang-bootstrap/utils/errors"
	"github.com/waldin-exe/golang-bootstrap/utils/response"
)

type UserHandler struct {
	service contract.UserService
}

func NewUserHandler(service contract.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Mount(router fiber.Router) {
	route := router.Group("/user")
	route.Get("/", h.GetUsers)
	route.Post("/", h.CreateUser)
	route.Put("/", h.UpdateUser)
	route.Delete("/", h.DeleteUser)
	route.Post("/ubah-password", h.UbahPassword)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	var req entity.GetUserRequest
	if err := c.QueryParser(&req); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	data, totalData, err := h.service.GetUsers(c.Context(), req)
	if err != nil {
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	payload := map[string]interface{}{
		"rows":       data,
		"total_data": totalData,
	}

	return response.NewResponseSuccess(c, payload, "Success")
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req entity.PostUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	userID := fmt.Sprintf("%v", c.Locals("user_id"))
	req.CreatedBy = userID

	if err := h.service.CreateUser(c.Context(), req); err != nil {
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseCreated(c, nil, "User berhasil ditambahkan")
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var req entity.PutUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, "invalid request body")
	}

	userID := fmt.Sprintf("%v", c.Locals("user_id"))
	req.UpdatedBy = userID

	if req.Id == 0 {
		return response.NewResponseBadRequest(c, "id is required")
	}

	if err := h.service.UpdateUser(c.Context(), req); err != nil {
		if appErrors.IsNotFoundError(err) {
			return response.NewResponseError(c, fiber.StatusNotFound, err.Error())
		}
		if appErrors.IsBadRequestError(err) {
			return response.NewResponseBadRequest(c, err.Error())
		}
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseSuccess(c, nil, "User berhasil diperbarui")
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	var req entity.DeleteUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, "invalid request body")
	}

	userID := fmt.Sprintf("%v", c.Locals("user_id"))
	req.UpdatedBy = userID

	if err := h.service.DeleteUser(c.Context(), req); err != nil {
		if appErrors.IsNotFoundError(err) {
			return response.NewResponseError(c, fiber.StatusNotFound, err.Error())
		}
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseSuccess(c, nil, "User berhasil dihapus")
}

func (h *UserHandler) UbahPassword(c *fiber.Ctx) error {
	var req entity.UbahPasswordInput
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, "Input tidak valid")
	}

	if err := h.service.UbahPassword(c.Context(), req); err != nil {
		if appErrors.IsBadRequestError(err) {
			return response.NewResponseBadRequest(c, err.Error())
		}
		if appErrors.IsNotFoundError(err) {
			return response.NewResponseError(c, fiber.StatusNotFound, err.Error())
		}
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseSuccess(c, nil, "Password berhasil diubah")
}
