package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/armada/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/armada/entity"
	gambarContract "github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/contract"
	gambarEntity "github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/entity"
	appErrors "github.com/waldin-exe/golang-bootstrap/utils/errors"
)

type armadaService struct {
	repo          contract.ArmadaRepository
	gambarService gambarContract.GambarService
}

func NewArmadaService(repo contract.ArmadaRepository, gambarService gambarContract.GambarService) contract.ArmadaService {
	return &armadaService{
		repo:          repo,
		gambarService: gambarService,
	}
}

func (s *armadaService) GetArmadas(ctx context.Context, req entity.GetArmadaRequest) ([]entity.ArmadaResponse, int64, error) {
	data, count, err := s.repo.GetArmadas(ctx, req)
	if err != nil {
		return nil, 0, appErrors.NewInternalError(err.Error())
	}
	if data == nil {
		data = []entity.ArmadaResponse{}
	}

	for i := range data {
		images, err := s.gambarService.GetGambarByReference(ctx, gambarEntity.GetGambarRequest{
			Reference: "armada",
			ForeignID: data[i].ID,
		})
		if err == nil {
			data[i].Images = images
		} else {
			data[i].Images = []gambarEntity.Gambar{}
		}
	}

	return data, count, nil
}

func (s *armadaService) GetArmadasTersedia(ctx context.Context, req entity.GetArmadaTersediaRequest) ([]entity.ArmadaResponse, int64, error) {
	data, count, err := s.repo.GetArmadasTersedia(ctx, req)
	if err != nil {
		return nil, 0, appErrors.NewInternalError(err.Error())
	}
	if data == nil {
		data = []entity.ArmadaResponse{}
	}

	for i := range data {
		images, err := s.gambarService.GetGambarByReference(ctx, gambarEntity.GetGambarRequest{
			Reference: "armada",
			ForeignID: data[i].ID,
		})
		if err == nil {
			data[i].Images = images
		} else {
			data[i].Images = []gambarEntity.Gambar{}
		}
	}

	return data, count, nil
}

func (s *armadaService) CreateArmada(ctx context.Context, req entity.PostArmadaRequest, files []*multipart.FileHeader) error {
	armada := &entity.Armada{
		PlatNomor:     req.PlatNomor,
		NomorLambung:  req.NomorLambung,
		JumlahSeat:    req.JumlahSeat,
		Merk:          req.Merk,
		Tahun:         req.Tahun,
		NoKIR:         req.NoKIR,
		IdJenisArmada: req.JenisNum,
		Body:          req.Body,
		CreatedBy:     req.CreatedBy,
	}

	if req.MasaBerlakuKIR != nil && *req.MasaBerlakuKIR != "" {
		parsedTime, err := time.Parse("2006-01-02", *req.MasaBerlakuKIR)
		if err != nil {
			return appErrors.NewBadRequestError("Format tanggal MasaBerlakuKIR tidak valid (YYYY-MM-DD)")
		}
		armada.MasaBerlakuKIR = &parsedTime
	}

	if err := s.repo.CreateArmada(ctx, armada); err != nil {
		return appErrors.NewInternalError(err.Error())
	}

	if len(files) > 0 {
		if err := s.gambarService.UploadImages(ctx, "armada", armada.ID, files); err != nil {
			// Rollback Armada creation manually
			_ = s.repo.DeleteArmada(ctx, armada.ID, "system_rollback")
			return err
		}
	}

	return nil
}

func (s *armadaService) UpdateArmada(ctx context.Context, req entity.PutArmadaRequest, files []*multipart.FileHeader, deletedImageIDs []int) error {
	armada, err := s.repo.FindByID(ctx, req.Id)
	if err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	if armada == nil {
		return appErrors.NewNotFoundError("Armada not found")
	}

	armada.PlatNomor = req.PlatNomor
	armada.NomorLambung = req.NomorLambung
	armada.JumlahSeat = req.JumlahSeat
	armada.Merk = req.Merk
	armada.Tahun = req.Tahun
	armada.NoKIR = req.NoKIR
	armada.IdJenisArmada = req.JenisNum
	armada.Body = req.Body
	armada.UpdatedBy = req.UpdatedBy

	if err := s.repo.UpdateArmada(ctx, armada); err != nil {
		return appErrors.NewInternalError(err.Error())
	}

	if len(deletedImageIDs) > 0 {
		_ = s.gambarService.DeleteGambar(ctx, gambarEntity.DeleteGambarRequest{
			DeleteImageIds: deletedImageIDs,
		})
	}

	if len(files) > 0 {
		if err := s.gambarService.UploadImages(ctx, "armada", armada.ID, files); err != nil {
			return err
		}
	}

	return nil
}

func (s *armadaService) DeleteArmada(ctx context.Context, req entity.DeleteArmadaRequest) error {
	if err := s.repo.DeleteArmada(ctx, req.Id, req.UpdatedBy); err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	return nil
}
