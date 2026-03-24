package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/auth/handler"
)

type AuthRoutes struct {
	Router      fiber.Router
	AuthHandler *handler.AuthHandler
}

func NewAuthRoutes(
	router fiber.Router,
	handler *handler.AuthHandler,
) *AuthRoutes {
	return &AuthRoutes{
		Router:      router,
		AuthHandler: handler,
	}
}

func (r *AuthRoutes) SetupRoutes() {
	group := r.Router.Group("/auth")

	group.Post("/login",
		r.AuthHandler.Login,
	)

	group.Post("/refresh",
		r.AuthHandler.RefreshToken,
	)
}
