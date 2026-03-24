package contract

import (
	"context"

	"gorm.io/gorm"
)

type DBProvider interface {
	GetDB() *gorm.DB
	WithTx(ctx context.Context, fn func(tx *gorm.DB) error) error
}
