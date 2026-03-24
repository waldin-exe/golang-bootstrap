package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/pegawai/entity"
	"gorm.io/gorm"
)

type pegawaiRepository struct {
	db *gorm.DB
}

func NewPegawaiRepository(db *gorm.DB) contract.PegawaiRepository {
	return &pegawaiRepository{db: db}
}

func (r *pegawaiRepository) GetPegawais(ctx context.Context, filter entity.GetPegawaiRequest) ([]entity.PegawaiResponse, int64, error) {
	var pegawais []entity.PegawaiResponse
	var totalData int64

	query := r.db.WithContext(ctx).Table("pegawais").
		Select("id, nama as nama_pegawai, tgl_lahir, alamat, no_telepon, jenis_pegawai").
		Where("deleted_at IS NULL")

	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.NamaPegawai != "" {
		query = query.Where("nama ILIKE ?", "%"+filter.NamaPegawai+"%")
	}
	if filter.JenisPegawai != "" {
		query = query.Where("jenis_pegawai ILIKE ?", "%"+filter.JenisPegawai+"%")
	}

	if filter.TanggalBerangkat != "" && filter.TanggalKembali != "" {
		notInClause := ""
		if len(filter.ExcludeIDs) > 0 {
			idsStr := []string{}
			for _, id := range filter.ExcludeIDs {
				idsStr = append(idsStr, fmt.Sprintf("%d", id))
			}
			notInClause = fmt.Sprintf("AND id_driver NOT IN (%s)", strings.Join(idsStr, ","))
			notInClause2 := strings.Replace(notInClause, "id_driver", "id_driver2", 1)
			notInClause3 := strings.Replace(notInClause, "id_driver", "id_kondektur", 1)

			subQuery := fmt.Sprintf(`
				SELECT DISTINCT id_driver FROM spjs 
				WHERE tanggal_berangkat <= ? AND tanggal_kembali >= ? AND deleted_at IS NULL %s
				UNION
				SELECT DISTINCT id_driver2 FROM spjs 
				WHERE tanggal_berangkat <= ? AND tanggal_kembali >= ? AND deleted_at IS NULL %s
				UNION
				SELECT DISTINCT id_kondektur FROM spjs 
				WHERE tanggal_berangkat <= ? AND tanggal_kembali >= ? AND deleted_at IS NULL %s
			`, notInClause, notInClause2, notInClause3)

			query = query.Where(fmt.Sprintf("id NOT IN (%s)", subQuery),
				filter.TanggalKembali, filter.TanggalBerangkat,
				filter.TanggalKembali, filter.TanggalBerangkat,
				filter.TanggalKembali, filter.TanggalBerangkat,
			)
		} else {
			subQuery := `
				SELECT DISTINCT id_driver FROM spjs 
				WHERE tanggal_berangkat <= ? AND tanggal_kembali >= ? AND deleted_at IS NULL
				UNION
				SELECT DISTINCT id_driver2 FROM spjs 
				WHERE tanggal_berangkat <= ? AND tanggal_kembali >= ? AND deleted_at IS NULL
				UNION
				SELECT DISTINCT id_kondektur FROM spjs 
				WHERE tanggal_berangkat <= ? AND tanggal_kembali >= ? AND deleted_at IS NULL
			`
			query = query.Where(fmt.Sprintf("id NOT IN (%s)", subQuery),
				filter.TanggalKembali, filter.TanggalBerangkat,
				filter.TanggalKembali, filter.TanggalBerangkat,
				filter.TanggalKembali, filter.TanggalBerangkat,
			)
		}
	}

	err := query.Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	limit := filter.Limit
	offset := filter.Offset
	if limit == 0 {
		limit = 10
	}

	err = query.Limit(limit).Offset(offset).Find(&pegawais).Error
	if err != nil {
		return nil, 0, err
	}

	return pegawais, totalData, nil
}

func (r *pegawaiRepository) FindByID(ctx context.Context, id int) (*entity.Pegawai, error) {
	var pegawai entity.Pegawai
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&pegawai).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &pegawai, nil
}

func (r *pegawaiRepository) CreatePegawai(ctx context.Context, pegawai *entity.Pegawai) error {
	return r.db.WithContext(ctx).Create(pegawai).Error
}

func (r *pegawaiRepository) UpdatePegawai(ctx context.Context, pegawai *entity.Pegawai) error {
	return r.db.WithContext(ctx).Save(pegawai).Error
}

func (r *pegawaiRepository) DeletePegawai(ctx context.Context, id int, updatedBy string) error {
	result := r.db.WithContext(ctx).Model(&entity.Pegawai{}).Where("id = ?", id).Updates(map[string]interface{}{
		"updated_by": updatedBy,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("pegawai not found")
	}
	return r.db.WithContext(ctx).Delete(&entity.Pegawai{}, id).Error
}
