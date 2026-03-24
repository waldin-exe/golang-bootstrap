package main

import (
	"log"

	"github.com/waldin-exe/golang-bootstrap/config"
	"github.com/waldin-exe/golang-bootstrap/internal/infrastructure/database/migration"
)

func main() {
	cfg := config.LoadConfig()

	log.Println("[CMD] running migration...")

	migrator := migration.NewMigrator(
		cfg.Database.URL(),
		"file://internal/infrastructure/database/migration/tables",
	)

	migrator.Run()

	log.Println("✅ migration finished")
}
