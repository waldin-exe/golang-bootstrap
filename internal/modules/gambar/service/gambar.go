package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/entity"
	appErrors "github.com/waldin-exe/golang-bootstrap/utils/errors"
)

type gambarService struct {
	repo contract.GambarRepository
}

func NewGambarService(repo contract.GambarRepository) contract.GambarService {
	return &gambarService{repo: repo}
}

func (s *gambarService) GetGambarByReference(ctx context.Context, req entity.GetGambarRequest) ([]entity.Gambar, error) {
	gambars, err := s.repo.GetGambarByReference(ctx, req.Reference, req.ForeignID)
	if err != nil {
		return nil, appErrors.NewInternalError(err.Error())
	}
	if gambars == nil {
		gambars = []entity.Gambar{}
	}
	return gambars, nil
}

func (s *gambarService) DeleteGambar(ctx context.Context, req entity.DeleteGambarRequest) error {
	for _, id := range req.DeleteImageIds {
		if err := s.repo.DeleteGambarByID(ctx, id); err != nil {
			return appErrors.NewInternalError(err.Error())
		}
	}
	return nil
}

func (s *gambarService) UploadImages(ctx context.Context, reference string, foreignID int, files []*multipart.FileHeader) error {
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".pdf":  true,
	}

	allowedMimeTypes := map[string]bool{
		"image/jpeg":      true,
		"image/png":       true,
		"application/pdf": true,
	}

	const maxFileSize = 5 * 1024 * 1024

	uploadDir := fmt.Sprintf("./uploads/%s", reference)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return appErrors.NewInternalError("Gagal membuat direktori penyimpanan: " + err.Error())
	}

	for _, file := range files {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !allowedExtensions[ext] {
			return appErrors.NewBadRequestError(fmt.Sprintf("File %s memiliki ekstensi tidak diperbolehkan", file.Filename))
		}

		contentType := file.Header.Get("Content-Type")
		if !allowedMimeTypes[contentType] {
			return appErrors.NewBadRequestError(fmt.Sprintf("File %s memiliki tipe file tidak diperbolehkan", file.Filename))
		}

		if file.Size > maxFileSize {
			return appErrors.NewBadRequestError(fmt.Sprintf("File %s melebihi ukuran maksimum 5MB", file.Filename))
		}

		filename := fmt.Sprintf("%s_%d_%s", reference, foreignID, file.Filename)
		savePath := fmt.Sprintf("%s/%s", uploadDir, filename)
		imgPath := fmt.Sprintf("/uploads/%s/%s", reference, filename)

		// Save file to disk
		src, err := file.Open()
		if err != nil {
			return appErrors.NewInternalError(fmt.Sprintf("Gagal membuka file %s: %s", file.Filename, err.Error()))
		}
		defer src.Close()

		dst, err := os.Create(savePath)
		if err != nil {
			return appErrors.NewInternalError(fmt.Sprintf("Gagal membuat file tujuan %s: %s", file.Filename, err.Error()))
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return appErrors.NewInternalError(fmt.Sprintf("Gagal menyalin file %s: %s", file.Filename, err.Error()))
		}

		// Save to DB
		g := &entity.Gambar{
			Reference:    reference,
			ForeignID:    foreignID,
			Path:         imgPath,
			OriginalName: file.Filename,
			MimeType:     file.Header.Get("Content-Type"),
		}
		if err := s.repo.CreateGambar(ctx, g); err != nil {
			return appErrors.NewInternalError(fmt.Sprintf("Gagal menyimpan metadata gambar %s: %s", file.Filename, err.Error()))
		}
	}

	return nil
}
