package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pegawai struct {
	ID           int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama         string         `gorm:"type:varchar(100);not null" json:"nama"`
	TglLahir     string         `gorm:"type:varchar(20)" json:"tgl_lahir"`
	Alamat       string         `gorm:"type:text" json:"alamat"`
	NoTelepon    string         `gorm:"type:varchar(20)" json:"no_telepon"`
	JenisPegawai string         `gorm:"type:varchar(50)" json:"jenis_pegawai"`
	CreatedBy    string         `gorm:"type:varchar(100)" json:"created_by"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedBy    string         `gorm:"type:varchar(100)" json:"updated_by"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Pegawai) TableName() string {
	return "pegawais"
}

type GetPegawaiRequest struct {
	Id               int    `query:"id"`
	NamaPegawai      string `query:"nama_pegawai"`
	JenisPegawai     string `query:"jenis_pegawai"`
	TanggalBerangkat string `query:"tanggal_berangkat"`
	TanggalKembali   string `query:"tanggal_kembali"`
	ExcludeIDs       []int  `query:"exclude_ids"`
	Limit            int    `query:"limit"`
	Offset           int    `query:"offset"`
}

type PostPegawaiRequest struct {
	NamaPegawai  string `json:"nama_pegawai" validate:"required"`
	TglLahir     string `json:"tgl_lahir,omitempty"`
	Alamat       string `json:"alamat" validate:"required"`
	NoTelepon    string `json:"no_telepon" validate:"required"`
	Level        string `json:"level,omitempty"`
	JenisPegawai string `json:"jenis_pegawai,omitempty"`
	CreatedBy    string `json:"created_by" `
}

type PutPegawaiRequest struct {
	Id           int    `json:"id" validate:"required"`
	NamaPegawai  string `json:"nama_pegawai" validate:"required"`
	TglLahir     string `json:"tgl_lahir,omitempty"`
	Alamat       string `json:"alamat" validate:"required"`
	Level        string `json:"level,omitempty"`
	JenisPegawai string `json:"jenis_pegawai,omitempty"`
	NoTelepon    string `json:"no_telepon" validate:"required"`
	UpdatedBy    string `json:"updated_by" `
}

type DeletePegawaiRequest struct {
	Id        int    `json:"id" validate:"required"`
	UpdatedBy string `json:"updated_by" `
}

type PegawaiResponse struct {
	Id           int    `json:"id"`
	NamaPegawai  string `json:"nama_pegawai"`
	TglLahir     string `json:"tgl_lahir"`
	Alamat       string `json:"alamat"`
	NoTelepon    string `json:"no_telepon"`
	JenisPegawai string `json:"jenis_pegawai"`
}
