package repository

import (
	"context"
	"errors"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/armada/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/armada/entity"
	"gorm.io/gorm"
)

type armadaRepository struct {
	db *gorm.DB
}

func NewArmadaRepository(db *gorm.DB) contract.ArmadaRepository {
	return &armadaRepository{db: db}
}

func (r *armadaRepository) GetArmadas(ctx context.Context, filter entity.GetArmadaRequest) ([]entity.ArmadaResponse, int64, error) {
	var armadas []entity.ArmadaResponse
	var totalData int64

	query := r.db.WithContext(ctx).Table("armadas as a").
		Select("a.id, a.plat_nomor, a.nomor_lambung, a.jumlah_seat, a.merk, a.tahun, a.no_kir, a.masa_berlaku_kir, ja.id as id_jenis_armada, ja.nama as jenis, a.body").
		Joins("JOIN jenis_armadas as ja ON ja.id = a.id_jenis_armada").
		Where("a.deleted_at IS NULL")

	if filter.Id != 0 {
		query = query.Where("a.id = ?", filter.Id)
	}
	if filter.PlatNomor != "" {
		query = query.Where("a.plat_nomor ILIKE ?", "%"+filter.PlatNomor+"%")
	}
	if filter.Merk != "" {
		query = query.Where("a.merk ILIKE ?", "%"+filter.Merk+"%")
	}
	if filter.Jenis != "" {
		query = query.Where("ja.id = ?", filter.Jenis)
	}
	if filter.JumlahSeat != 0 {
		query = query.Where("a.jumlah_seat = ?", filter.JumlahSeat)
	}

	if err := query.Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	limit := filter.Limit
	offset := filter.Offset
	if limit == 0 {
		limit = 10
	}

	if err := query.Limit(limit).Offset(offset).Find(&armadas).Error; err != nil {
		return nil, 0, err
	}

	return armadas, totalData, nil
}

func (r *armadaRepository) GetArmadasTersedia(ctx context.Context, filter entity.GetArmadaTersediaRequest) ([]entity.ArmadaResponse, int64, error) {
	var armadas []entity.ArmadaResponse
	var totalData int64

	subQuery := `
		SELECT 1
		FROM jadwal_armadas j
		WHERE j.id_armada = a.id
		  AND j.tanggal_mulai <= ?
		  AND j.tanggal_selesai >= ?
		  AND j.deleted_at IS NULL
	`

	query := r.db.WithContext(ctx).Table("armadas as a").
		Select("a.id, a.plat_nomor, a.nomor_lambung, a.jumlah_seat, a.merk, a.tahun, a.no_kir, a.masa_berlaku_kir, ja.id as id_jenis_armada, ja.nama as jenis, a.body").
		Joins("JOIN jenis_armadas as ja ON ja.id = a.id_jenis_armada").
		Where("NOT EXISTS ("+subQuery+")", filter.TanggalAkhir, filter.TanggalMulai).
		Where("a.deleted_at IS NULL")

	if err := query.Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	limit := filter.Limit
	offset := filter.Offset
	if limit == 0 {
		limit = 10
	}

	if err := query.Limit(limit).Offset(offset).Find(&armadas).Error; err != nil {
		return nil, 0, err
	}

	return armadas, totalData, nil
}

func (r *armadaRepository) GetArmadaByID(ctx context.Context, id int) (*entity.ArmadaResponse, error) {
	var armada entity.ArmadaResponse
	query := r.db.WithContext(ctx).Table("armadas as a").
		Select("a.id, a.plat_nomor, a.nomor_lambung, a.jumlah_seat, a.merk, a.tahun, a.no_kir, a.masa_berlaku_kir, ja.id as id_jenis_armada, ja.nama as jenis, a.body").
		Joins("JOIN jenis_armadas as ja ON ja.id = a.id_jenis_armada").
		Where("a.id = ?", id)

	if err := query.First(&armada).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &armada, nil
}

func (r *armadaRepository) CreateArmada(ctx context.Context, armada *entity.Armada) error {
	return r.db.WithContext(ctx).Create(armada).Error
}

func (r *armadaRepository) UpdateArmada(ctx context.Context, armada *entity.Armada) error {
	return r.db.WithContext(ctx).Save(armada).Error
}

func (r *armadaRepository) DeleteArmada(ctx context.Context, id int, updatedBy string) error {
	result := r.db.WithContext(ctx).Model(&entity.Armada{}).Where("id = ?", id).Updates(map[string]interface{}{
		"updated_by": updatedBy,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("armada not found")
	}
	return r.db.WithContext(ctx).Delete(&entity.Armada{}, id).Error
}

func (r *armadaRepository) FindByID(ctx context.Context, id int) (*entity.Armada, error) {
	var armada entity.Armada
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&armada).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &armada, nil
}
