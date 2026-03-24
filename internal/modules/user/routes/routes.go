package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/user/handler"
	"github.com/waldin-exe/golang-bootstrap/internal/shared/contract"
)

type UserRoutes struct {
	Router            fiber.Router
	UserHandler       *handler.UserHandler
	MiddlewareManager contract.MiddlewareManager
}

func NewUserRoutes(
	router fiber.Router,
	handler *handler.UserHandler,
	middleware contract.MiddlewareManager,
) *UserRoutes {
	return &UserRoutes{
		Router:            router,
		UserHandler:       handler,
		MiddlewareManager: middleware,
	}
}

func (r *UserRoutes) SetupRoutes() {
	group := r.Router.Group(
		"/user",
	)

	group.Get("/",
		r.MiddlewareManager.RequireRoles("admin", "owner"),
		r.UserHandler.GetUsers,
	)
	group.Post("/add",
		r.MiddlewareManager.RequireRoles("admin", "owner"),
		r.UserHandler.CreateUser,
	)
	group.Put("/update",
		r.MiddlewareManager.RequireRoles("admin", "owner"),
		r.UserHandler.UpdateUser,
	)
	group.Delete("/delete",
		r.MiddlewareManager.RequireRoles("admin", "owner"),
		r.UserHandler.DeleteUser,
	)
	group.Post("/ubah-password",
		r.UserHandler.UbahPassword,
	)

}
