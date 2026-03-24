package contract

import (
	"context"
	"mime/multipart"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/entity"
)

type GambarRepository interface {
	GetGambarByReference(ctx context.Context, reference string, foreignID int) ([]entity.Gambar, error)
	CreateGambar(ctx context.Context, gambar *entity.Gambar) error
	DeleteGambarByID(ctx context.Context, id int) error
}

type GambarService interface {
	GetGambarByReference(ctx context.Context, req entity.GetGambarRequest) ([]entity.Gambar, error)
	DeleteGambar(ctx context.Context, req entity.DeleteGambarRequest) error
	UploadImages(ctx context.Context, reference string, foreignID int, files []*multipart.FileHeader) error
}
