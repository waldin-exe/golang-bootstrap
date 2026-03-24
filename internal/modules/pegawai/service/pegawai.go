package service

import (
	"context"
	"strings"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/entity"
	userContract "github.com/waldin-exe/golang-bootstrap/internal/modules/user/contract"
	userEntity "github.com/waldin-exe/golang-bootstrap/internal/modules/user/entity"
	appErrors "github.com/waldin-exe/golang-bootstrap/utils/errors"
)

type pegawaiService struct {
	repo        contract.PegawaiRepository
	userService userContract.UserService
}

func NewPegawaiService(repo contract.PegawaiRepository, userService userContract.UserService) contract.PegawaiService {
	return &pegawaiService{
		repo:        repo,
		userService: userService,
	}
}

func (s *pegawaiService) GetPegawais(ctx context.Context, req entity.GetPegawaiRequest) ([]entity.PegawaiResponse, int64, error) {
	pegawais, count, err := s.repo.GetPegawais(ctx, req)
	if err != nil {
		return nil, 0, appErrors.NewInternalError(err.Error())
	}
	if pegawais == nil {
		pegawais = []entity.PegawaiResponse{}
	}
	return pegawais, count, nil
}

func (s *pegawaiService) CreatePegawai(ctx context.Context, req entity.PostPegawaiRequest) error {
	if req.Level == "" {
		req.Level = req.JenisPegawai
	}

	pegawai := &entity.Pegawai{
		Nama:         req.NamaPegawai,
		TglLahir:     req.TglLahir,
		Alamat:       req.Alamat,
		NoTelepon:    req.NoTelepon,
		JenisPegawai: req.JenisPegawai,
		CreatedBy:    req.CreatedBy,
		UpdatedBy:    req.CreatedBy,
	}

	if err := s.repo.CreatePegawai(ctx, pegawai); err != nil {
		return appErrors.NewInternalError(err.Error())
	}

	// Create User account automatically
	userReq := userEntity.PostUserRequest{
		Username:  strings.ToLower(strings.ReplaceAll(req.NamaPegawai, " ", "")),
		Password:  "usergenerated",
		Level:     req.Level,
		PegawaiId: pegawai.ID,
		CreatedBy: req.CreatedBy,
	}

	if err := s.userService.CreateUser(ctx, userReq); err != nil {
		return appErrors.NewInternalError("Pegawai created but failed to create user account: " + err.Error())
	}

	return nil
}

func (s *pegawaiService) UpdatePegawai(ctx context.Context, req entity.PutPegawaiRequest) error {
	pegawai, err := s.repo.FindByID(ctx, req.Id)
	if err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	if pegawai == nil {
		return appErrors.NewNotFoundError("Pegawai not found")
	}

	if req.Level == "" {
		req.Level = req.JenisPegawai
	}

	pegawai.Nama = req.NamaPegawai
	pegawai.TglLahir = req.TglLahir
	pegawai.Alamat = req.Alamat
	pegawai.NoTelepon = req.NoTelepon
	pegawai.JenisPegawai = req.JenisPegawai
	pegawai.UpdatedBy = req.UpdatedBy

	if err := s.repo.UpdatePegawai(ctx, pegawai); err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	return nil
}

func (s *pegawaiService) DeletePegawai(ctx context.Context, req entity.DeletePegawaiRequest) error {
	pegawai, err := s.repo.FindByID(ctx, req.Id)
	if err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	if pegawai == nil {
		return appErrors.NewNotFoundError("Pegawai not found")
	}

	if err := s.repo.DeletePegawai(ctx, req.Id, req.UpdatedBy); err != nil {
		return appErrors.NewInternalError(err.Error())
	}

	return nil
}
