package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponseSuccess(c *fiber.Ctx, data any, message string) error {
	if message == "" {
		message = "success"
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  fiber.StatusOK,
		Message: message,
		Data:    data,
	})
}

func NewResponseCreated(c *fiber.Ctx, data interface{}, message string) error {
	if message == "" {
		message = "created"
	}
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status:  fiber.StatusCreated,
		Message: message,
		Data:    data,
	})
}

func NewResponseBadRequest(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "bad_request"
	}
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:  fiber.StatusBadRequest,
		Message: message,
	})
}

func NewResponseUnauthorized(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "unauthorized"
	}
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Status:  fiber.StatusUnauthorized,
		Message: message,
	})
}

func NewResponseForbidden(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "forbidden"
	}
	return c.Status(fiber.StatusForbidden).JSON(Response{
		Status:  fiber.StatusForbidden,
		Message: message,
	})
}

func NewResponseError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Response{
		Status:  status,
		Message: message,
	})
}
