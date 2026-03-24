package entity

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	UserEmail string         `gorm:"type:varchar(100);not null;index" json:"user_email"`
	Token     string         `gorm:"type:text;not null" json:"token"`
	ExpiredAt time.Time      `gorm:"not null" json:"expired_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type UserAuthResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Nama      string `json:"nama"`
	TglLahir  string `json:"tgl_lahir"`
	Alamat    string `json:"alamat"`
	NoTelepon string `json:"no_telepon"`
}

type TokenResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int64     `json:"expires_in"`
	ExpiresDate  time.Time `json:"expires_date"`
}

type AuthUserClaims struct {
	UserID    int
	Username  string
	PegawaiID int
	Role      string
}
