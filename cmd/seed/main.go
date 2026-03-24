package main

import (
	"context"
	"log"

	"github.com/waldin-exe/golang-bootstrap/config"
	"github.com/waldin-exe/golang-bootstrap/internal/infrastructure/database"
	"github.com/waldin-exe/golang-bootstrap/internal/infrastructure/database/migration/seed"
)

func main() {
	cfg := config.LoadConfig()
	ctx := context.Background()

	log.Println("[CMD] running seeder...")

	// init DB
	db, err := database.NewGormDB(cfg.Database)
	if err != nil {
		log.Fatalf("failed connect DB: %v", err)
	}

	// init tx manager
	txManager := database.NewDBProvider(db)

	// init seeder
	userSeeder := seed.NewUserSeeder(txManager)

	// runner
	seedRunner := seed.NewRunner(
		userSeeder,
	)

	seedRunner.Run(ctx)

	log.Println("✅ seeding finished")
}
