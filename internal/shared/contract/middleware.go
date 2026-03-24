package contract

import "github.com/gofiber/fiber/v2"

type MiddlewareManager interface {
	Protected() fiber.Handler
	RequireRoles(roles ...string) fiber.Handler
}
