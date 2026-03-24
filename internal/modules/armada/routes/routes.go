package armada

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/armada/handler"
	"github.com/waldin-exe/golang-bootstrap/internal/shared/contract"
)

type ArmadaRoutes struct {
	Router            fiber.Router
	ArmadaHandler     *handler.ArmadaHandler
	MiddlewareManager contract.MiddlewareManager
}

func NewArmadaRoutes(
	router fiber.Router,
	handler *handler.ArmadaHandler,
	middleware contract.MiddlewareManager,
) *ArmadaRoutes {
	return &ArmadaRoutes{
		Router:            router,
		ArmadaHandler:     handler,
		MiddlewareManager: middleware,
	}
}

func (r *ArmadaRoutes) SetupRoutes() {
	group := r.Router.Group(
		"/armada",
		r.MiddlewareManager.Protected(),
		r.MiddlewareManager.RequireRoles("admin", "owner"),
	)

	group.Get("/", r.ArmadaHandler.GetArmadas)
	group.Post("/add", r.ArmadaHandler.CreateArmada)
	group.Put("/update", r.ArmadaHandler.UpdateArmada)
	group.Delete("/delete", r.ArmadaHandler.DeleteArmada)
	group.Get("/tersedia", r.ArmadaHandler.GetArmadasTersedia)
}
