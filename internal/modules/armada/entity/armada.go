package entity

import (
	"time"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/gambar/entity"
	"gorm.io/gorm"
)

type Armada struct {
	ID             int            `gorm:"primaryKey;autoIncrement" json:"id"`
	PlatNomor      string         `gorm:"type:varchar(50);not null" json:"plat_nomor"`
	NomorLambung   string         `gorm:"type:varchar(50);not null" json:"nomor_lambung"`
	JumlahSeat     int            `gorm:"not null" json:"jumlah_seat"`
	Merk           string         `gorm:"type:varchar(100);not null" json:"merk"`
	Tahun          string         `gorm:"type:varchar(4)" json:"tahun"`
	NoKIR          string         `gorm:"type:varchar(50)" json:"no_kir"`
	MasaBerlakuKIR *time.Time     `gorm:"type:date" json:"masa_berlaku_kir"`
	IdJenisArmada  int            `gorm:"not null" json:"id_jenis_armada"`
	Body           string         `gorm:"type:varchar(100)" json:"body"`
	CreatedBy      string         `gorm:"type:varchar(100)" json:"created_by"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedBy      string         `gorm:"type:varchar(100)" json:"updated_by"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

type GetArmadaRequest struct {
	Id         int    `query:"id"`
	PlatNomor  string `query:"plat_nomor"`
	Merk       string `query:"merk"`
	JumlahSeat int    `query:"jumlah_seat"`
	Jenis      string `query:"jenis"`
	Limit      int    `query:"limit"`
	Offset     int    `query:"offset"`
}

type GetArmadaTersediaRequest struct {
	TanggalMulai string `query:"tgl_mulai"`
	TanggalAkhir string `query:"tgl_akhir"`
	Limit        int    `query:"limit"`
	Offset       int    `query:"offset"`
}

type PostArmadaRequest struct {
	PlatNomor      string  `json:"plat_nomor" form:"plat_nomor" validate:"required"`
	NomorLambung   string  `json:"nomor_lambung" form:"nomor_lambung" validate:"required"`
	JumlahSeat     int     `json:"jumlah_seat" form:"jumlah_seat" validate:"required"`
	Merk           string  `json:"merk" form:"merk" validate:"required"`
	Tahun          string  `json:"tahun" form:"tahun"`
	NoKIR          string  `json:"no_kir" form:"no_kir"`
	MasaBerlakuKIR *string `json:"masa_berlaku_kir" form:"masa_berlaku_kir"`
	Jenis          string  `json:"jenis" form:"jenis" validate:"required"`
	JenisNum       int     `json:"jenis_number" form:"jenis_number"`
	Body           string  `json:"body" form:"body" validate:"required"`
	CreatedBy      string  `json:"created_by" form:"created_by" `
}

type PutArmadaRequest struct {
	Id             int     `json:"id" form:"id" validate:"required"`
	PlatNomor      string  `json:"plat_nomor" form:"plat_nomor" validate:"required"`
	NomorLambung   string  `json:"nomor_lambung" form:"nomor_lambung" validate:"required"`
	JumlahSeat     int     `json:"jumlah_seat" form:"jumlah_seat" validate:"required"`
	Merk           string  `json:"merk" form:"merk" validate:"required"`
	Tahun          string  `json:"tahun" form:"tahun"`
	NoKIR          string  `json:"no_kir" form:"no_kir"`
	MasaBerlakuKIR *string `json:"masa_berlaku_kir" form:"masa_berlaku_kir"`
	Jenis          string  `json:"jenis" form:"jenis" validate:"required"`
	JenisNum       int     `json:"jenis_number" form:"jenis_number"`
	Body           string  `json:"body" form:"body" validate:"required"`
	UpdatedBy      string  `json:"updated_by" form:"updated_by" `
}

type DeleteArmadaRequest struct {
	Id        int    `json:"id" validate:"required"`
	UpdatedBy string `json:"updated_by" `
}

type ArmadaResponse struct {
	ID             int             `json:"id"`
	PlatNomor      string          `json:"plat_nomor"`
	NomorLambung   string          `json:"nomor_lambung"`
	JumlahSeat     int             `json:"jumlah_seat"`
	Merk           string          `json:"merk"`
	Tahun          string          `json:"tahun"`
	NoKIR          string          `json:"no_kir"`
	MasaBerlakuKIR *time.Time      `json:"masa_berlaku_kir,omitempty"`
	IdJenisArmada  int             `json:"id_jenis_armada"`
	Jenis          string          `json:"jenis"`
	Body           string          `json:"body"`
	Images         []entity.Gambar `json:"images" gorm:"-"`
}
