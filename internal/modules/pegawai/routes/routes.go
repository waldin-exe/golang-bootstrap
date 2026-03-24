package pegawai

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/handler"
	"github.com/waldin-exe/golang-bootstrap/internal/shared/contract"
)

type PegawaiRoutes struct {
	Router            fiber.Router
	PegawaiHandler    *handler.PegawaiHandler
	MiddlewareManager contract.MiddlewareManager
}

func NewPegawaiRoutes(
	router fiber.Router,
	handler *handler.PegawaiHandler,
	middleware contract.MiddlewareManager,
) *PegawaiRoutes {
	return &PegawaiRoutes{
		Router:            router,
		PegawaiHandler:    handler,
		MiddlewareManager: middleware,
	}
}

func (r *PegawaiRoutes) SetupRoutes() {
	group := r.Router.Group(
		"/pegawai",
		r.MiddlewareManager.Protected(),
	)

	group.Get("/",
		r.PegawaiHandler.GetPegawais,
	)

	group.Post("/add",
		r.PegawaiHandler.CreatePegawai,
	)

	group.Put("/update",
		r.PegawaiHandler.UpdatePegawai,
	)

	group.Delete("/delete",
		r.PegawaiHandler.DeletePegawai,
	)

}
