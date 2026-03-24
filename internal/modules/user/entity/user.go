package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
	Level     string         `gorm:"type:varchar(50);not null" json:"level"`
	PegawaiID int            `gorm:"not null" json:"pegawai_id"`
	CreatedBy string         `gorm:"type:varchar(100)" json:"created_by"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedBy string         `gorm:"type:varchar(100)" json:"updated_by"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	DeletedBy *string        `gorm:"type:varchar(100)" json:"deleted_by"`
}

func (User) TableName() string {
	return "users"
}

type GetUserRequest struct {
	ID          int    `query:"id"`
	NamaPegawai string `query:"nama_pegawai"`
	Limit       int    `query:"limit"`
	Offset      int    `query:"offset"`
}

type PostUserRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Level     string `json:"level" validate:"required"`
	PegawaiId int    `json:"pegawai_id" validate:"required"`
	CreatedBy string `json:"created_by" `
}

type PutUserRequest struct {
	Id          int    `json:"id" validate:"required"`
	Username    string `json:"username" validate:"required"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	Level       string `json:"level" validate:"required"`
	UpdatedBy   string `json:"updated_by" `
}

type DeleteUserRequest struct {
	Id        int    `json:"id" validate:"required"`
	UpdatedBy string `json:"updated_by" `
}

type UserResponse struct {
	ID          int    `json:"id"`
	NamaPegawai string `json:"nama_pegawai"`
	Username    string `json:"username"`
	Level       string `json:"level"`
}

type UbahPasswordInput struct {
	UserID          int    `json:"user_id"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
