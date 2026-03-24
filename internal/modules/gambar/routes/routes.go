package gambar

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/handler"
	"github.com/waldin-exe/golang-bootstrap/internal/shared/contract"
)

type GambarRoutes struct {
	Router            fiber.Router
	GambarHandler     *handler.GambarHandler
	MiddlewareManager contract.MiddlewareManager
}

func NewGambarRoutes(
	router fiber.Router,
	handler *handler.GambarHandler,
	middleware contract.MiddlewareManager,
) *GambarRoutes {
	return &GambarRoutes{
		Router:            router,
		GambarHandler:     handler,
		MiddlewareManager: middleware,
	}
}

func (r *GambarRoutes) SetupRoutes() {
	group := r.Router.Group(
		"/gambar",
	)

	group.Get("/",
		r.MiddlewareManager.Protected(),
		r.GambarHandler.GetGambar,
	)

	group.Post("/add",
		r.MiddlewareManager.Protected(),
		r.GambarHandler.UploadGambar,
	)

	group.Delete("/delete",
		r.MiddlewareManager.Protected(),
		r.GambarHandler.DeleteGambar,
	)
}
