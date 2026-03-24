package setup

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// token
	"github.com/waldin-exe/golang-bootstrap/config"
	tokenService "github.com/waldin-exe/golang-bootstrap/internal/token"

	// auth
	authHandler "github.com/waldin-exe/golang-bootstrap/internal/modules/auth/handler"
	authRepo "github.com/waldin-exe/golang-bootstrap/internal/modules/auth/repository"
	authRoutes "github.com/waldin-exe/golang-bootstrap/internal/modules/auth/routes"
	authService "github.com/waldin-exe/golang-bootstrap/internal/modules/auth/service"

	// armada
	armadaHandler "github.com/waldin-exe/golang-bootstrap/internal/modules/armada/handler"
	armadaRepo "github.com/waldin-exe/golang-bootstrap/internal/modules/armada/repository"
	armadaRoutes "github.com/waldin-exe/golang-bootstrap/internal/modules/armada/routes"
	armadaService "github.com/waldin-exe/golang-bootstrap/internal/modules/armada/service"

	// gambar
	gambarHandler "github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/handler"
	gambarRepo "github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/repository"
	gambarRoutes "github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/routes"
	gambarService "github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/service"

	// pegawai
	pegawaiHandler "github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/handler"
	pegawaiRepo "github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/repository"
	pegawaiRoutes "github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/routes"
	pegawaiService "github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/service"

	// user
	userHandler "github.com/waldin-exe/golang-bootstrap/internal/modules/user/handler"
	userRepo "github.com/waldin-exe/golang-bootstrap/internal/modules/user/repository"
	userRoutes "github.com/waldin-exe/golang-bootstrap/internal/modules/user/routes"
	userService "github.com/waldin-exe/golang-bootstrap/internal/modules/user/service"
	"github.com/waldin-exe/golang-bootstrap/internal/shared/contract"
)

func InitModules(app *fiber.App, db *gorm.DB, middleware contract.MiddlewareManager, cfg *config.Config) {
	api := app.Group("/api")

	// ========================
	// AUTH MODULE
	// ========================

	// token
	tokenSvc := tokenService.NewTokenService(cfg.JWT)

	// repo
	authRepo := authRepo.NewAuthRepository(db)

	// service
	authSvc := authService.NewAuthService(authRepo, tokenSvc, *cfg)

	// handler
	authHandler := authHandler.NewAuthHandler(authSvc)

	// routes
	authRoutes := authRoutes.NewAuthRoutes(api, authHandler)
	authRoutes.SetupRoutes()

	// ========================
	// GAMBAR MODULE
	// ========================

	// repo
	gambarRepo := gambarRepo.NewGambarRepository(db)

	// service
	gambarSvc := gambarService.NewGambarService(gambarRepo)

	// handler
	gambarHandler := gambarHandler.NewGambarHandler(gambarSvc)

	// routes
	gambarRoutes := gambarRoutes.NewGambarRoutes(api, gambarHandler, middleware)
	gambarRoutes.SetupRoutes()

	// ========================
	// ARMADA MODULE
	// ========================

	// repo
	armadaRepo := armadaRepo.NewArmadaRepository(db)

	// service
	armadaSvc := armadaService.NewArmadaService(armadaRepo, gambarSvc)

	// handler
	armadaHandler := armadaHandler.NewArmadaHandler(armadaSvc)

	// routes
	armadaRoutes := armadaRoutes.NewArmadaRoutes(api, armadaHandler, middleware)
	armadaRoutes.SetupRoutes()

	// ========================
	// USER MODULE
	// ========================

	// repo
	userRepo := userRepo.NewUserRepository(db)

	// service
	userSvc := userService.NewUserService(userRepo)

	// handler
	userHandler := userHandler.NewUserHandler(userSvc)

	// routes
	userRoutes := userRoutes.NewUserRoutes(api, userHandler, middleware)
	userRoutes.SetupRoutes()

	// ========================
	// PEGAWAI MODULE
	// ========================

	// repo
	pegawaiRepo := pegawaiRepo.NewPegawaiRepository(db)

	// service
	pegawaiSvc := pegawaiService.NewPegawaiService(pegawaiRepo, userSvc)

	// handler
	pegawaiHandler := pegawaiHandler.NewPegawaiHandler(pegawaiSvc)

	// routes
	pegawaiRoutes := pegawaiRoutes.NewPegawaiRoutes(api, pegawaiHandler, middleware)
	pegawaiRoutes.SetupRoutes()

}
