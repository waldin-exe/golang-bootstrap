package database

import (
	"context"

	"github.com/waldin-exe/golang-bootstrap/internal/shared/contract"
	"gorm.io/gorm"
)

type dbProvider struct {
	db *gorm.DB
}

func NewDBProvider(db *gorm.DB) contract.DBProvider {
	return &dbProvider{db: db}
}

func (p *dbProvider) GetDB() *gorm.DB {
	return p.db
}

func (p *dbProvider) WithTx(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return p.db.WithContext(ctx).Transaction(fn)
}
