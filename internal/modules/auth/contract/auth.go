package contract

import (
	"context"
	"time"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/auth/entity"
)

// We define a struct to hold complete user info needed for claims
type AuthUser struct {
	Id        int
	Password  string
	Username  string
	Nama      string
	TglLahir  string
	Alamat    string
	NoTelepon string
	Level     string
	PegawaiID int
}

type AuthRepository interface {
	FindUserByUsername(ctx context.Context, username string) (*AuthUser, error)
	FindUserByEmail(ctx context.Context, email string) (*AuthUser, error)
	SaveRefreshToken(ctx context.Context, email, token string, expiredAt time.Time) error
	DeleteRefreshTokensByEmail(ctx context.Context, email string) error
	CheckRefreshTokenValid(ctx context.Context, email, token string) (bool, error)
	UpdateRefreshToken(ctx context.Context, email, token string, expiredAt time.Time) error
}

type AuthService interface {
	Login(ctx context.Context, req entity.LoginRequest) (*entity.TokenResponse, error)
	RefreshToken(ctx context.Context, req entity.RefreshTokenRequest) (*entity.TokenResponse, error)
}
