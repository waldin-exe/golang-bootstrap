package repository

import (
	"context"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/entity"
	"gorm.io/gorm"
)

type gambarRepository struct {
	db *gorm.DB
}

func NewGambarRepository(db *gorm.DB) contract.GambarRepository {
	return &gambarRepository{db: db}
}

func (r *gambarRepository) GetGambarByReference(ctx context.Context, reference string, foreignID int) ([]entity.Gambar, error) {
	var gambars []entity.Gambar
	err := r.db.WithContext(ctx).Where("reference = ? AND foreign_id = ?", reference, foreignID).Find(&gambars).Error
	return gambars, err
}

func (r *gambarRepository) CreateGambar(ctx context.Context, gambar *entity.Gambar) error {
	return r.db.WithContext(ctx).Create(gambar).Error
}

func (r *gambarRepository) DeleteGambarByID(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Gambar{}, id).Error
}
