package contract

import (
	"context"
	"mime/multipart"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/armada/entity"
)

type ArmadaRepository interface {
	GetArmadas(ctx context.Context, filter entity.GetArmadaRequest) ([]entity.ArmadaResponse, int64, error)
	GetArmadasTersedia(ctx context.Context, filter entity.GetArmadaTersediaRequest) ([]entity.ArmadaResponse, int64, error)
	GetArmadaByID(ctx context.Context, id int) (*entity.ArmadaResponse, error)
	CreateArmada(ctx context.Context, armada *entity.Armada) error
	UpdateArmada(ctx context.Context, armada *entity.Armada) error
	DeleteArmada(ctx context.Context, id int, updatedBy string) error
	FindByID(ctx context.Context, id int) (*entity.Armada, error)
}

type ArmadaService interface {
	GetArmadas(ctx context.Context, req entity.GetArmadaRequest) ([]entity.ArmadaResponse, int64, error)
	GetArmadasTersedia(ctx context.Context, req entity.GetArmadaTersediaRequest) ([]entity.ArmadaResponse, int64, error)
	CreateArmada(ctx context.Context, req entity.PostArmadaRequest, files []*multipart.FileHeader) error
	UpdateArmada(ctx context.Context, req entity.PutArmadaRequest, files []*multipart.FileHeader, deletedImageIDs []int) error
	DeleteArmada(ctx context.Context, req entity.DeleteArmadaRequest) error
}
