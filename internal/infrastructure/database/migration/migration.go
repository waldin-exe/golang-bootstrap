package migration

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
	DBUrl       string
	MigrateFile string
}

func NewMigrator(dbURL string, migrateFile string) *Migrator {
	return &Migrator{
		DBUrl:       dbURL,
		MigrateFile: migrateFile,
	}
}

func (m *Migrator) Run() {
	log.Println("[Database] running migration...")

	// 🔥 SQL Migration (golang-migrate)
	migrator, err := migrate.New(m.MigrateFile, m.DBUrl)
	if err != nil {
		log.Fatalf("failed create migrator: %v", err)
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("✅ SQL Migration done")
}
