package contract

import (
	"context"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/entity"
)

type PegawaiRepository interface {
	GetPegawais(ctx context.Context, filter entity.GetPegawaiRequest) ([]entity.PegawaiResponse, int64, error)
	CreatePegawai(ctx context.Context, pegawai *entity.Pegawai) error
	UpdatePegawai(ctx context.Context, pegawai *entity.Pegawai) error
	DeletePegawai(ctx context.Context, id int, updatedBy string) error
	FindByID(ctx context.Context, id int) (*entity.Pegawai, error)
}

type PegawaiService interface {
	GetPegawais(ctx context.Context, req entity.GetPegawaiRequest) ([]entity.PegawaiResponse, int64, error)
	CreatePegawai(ctx context.Context, req entity.PostPegawaiRequest) error
	UpdatePegawai(ctx context.Context, req entity.PutPegawaiRequest) error
	DeletePegawai(ctx context.Context, req entity.DeletePegawaiRequest) error
}
