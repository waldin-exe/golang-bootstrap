package database

import (
	"errors"
	"time"

	"github.com/waldin-exe/golang-bootstrap/config/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB(cfg database.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}

func Close(db *gorm.DB) error {
	if db == nil {
		return nil // no-op, aman
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if sqlDB == nil {
		return errors.New("sql.DB is nil")
	}

	return sqlDB.Close()
}
