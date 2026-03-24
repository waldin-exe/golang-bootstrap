package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/auth/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/auth/entity"
	appErrors "github.com/waldin-exe/golang-bootstrap/utils/errors"
	"github.com/waldin-exe/golang-bootstrap/utils/response"
)

type AuthHandler struct {
	service contract.AuthService
}

func NewAuthHandler(service contract.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req entity.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	token, err := h.service.Login(c.Context(), req)
	if err != nil {
		if appErrors.IsNotFoundError(err) || appErrors.IsBadRequestError(err) || appErrors.IsNotFoundError(err) {
			return response.NewResponseError(c, fiber.StatusUnauthorized, err.Error())
		}
		return response.NewResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.NewResponseSuccess(c, token, "Success")
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req entity.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil || req.RefreshToken == "" {
		return response.NewResponseBadRequest(c, "refresh_token is required")
	}

	token, err := h.service.RefreshToken(c.Context(), req)
	if err != nil {
		return response.NewResponseError(c, fiber.StatusUnauthorized, err.Error())
	}

	return response.NewResponseSuccess(c, token, "Success")
}
