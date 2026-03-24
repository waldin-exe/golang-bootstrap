package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/waldin-exe/golang-bootstrap/cmd/app/setup"
	"github.com/waldin-exe/golang-bootstrap/config"
	"github.com/waldin-exe/golang-bootstrap/internal/infrastructure/database"
	"github.com/waldin-exe/golang-bootstrap/middleware"
)

func main() {
	log.Println("[APP] starting application...")

	// Load config
	cfg := config.LoadConfig()

	// Init DB
	db, err := database.NewGormDB(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			log.Println("failed to close DB:", err)
		}
	}()

	// Fiber config
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   8 * 1024 * 1024,
	})

	// Middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://skripsi.sorakawa-wildan.my.id",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization",
		AllowCredentials: false,
	}))

	// middleware
	middlewareManager := middleware.NewMiddlewareManager(
		cfg.App.JWTSecret,
	)

	// Logger middleware
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
	}))

	// init semua module di sini
	setup.InitModules(app, db, middlewareManager, cfg)

	// Static file
	app.Static("/uploads", "./uploads")

	// Start server
	log.Printf("[APP] running on port %s", cfg.App.Port)
	if err := app.Listen(":" + cfg.App.Port); err != nil {
		log.Fatal(err)
	}
}
