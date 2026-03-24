package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	jwtConfig "github.com/waldin-exe/golang-bootstrap/config/jwt"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/auth/entity"
	"github.com/waldin-exe/golang-bootstrap/internal/token/contract"
)

type TokenService struct {
	cfg jwtConfig.JWTConfig
}

var _ contract.TokenService = (*TokenService)(nil)

func NewTokenService(cfg jwtConfig.JWTConfig) *TokenService {
	return &TokenService{cfg: cfg}
}

func (s *TokenService) GenerateToken(user entity.AuthUserClaims) (*entity.TokenResponse, error) {
	// ACCESS TOKEN
	accessExp := time.Now().Add(time.Minute * time.Duration(s.cfg.AccessExpireMinutes))

	accessClaims := jwt.MapClaims{
		"username":   user.Username,
		"user_id":    user.UserID,
		"pegawai_id": user.PegawaiID,
		"role":       user.Role,
		"exp":        accessExp.Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenStr, err := accessToken.SignedString([]byte(s.cfg.AccessSecret))
	if err != nil {
		return nil, err
	}

	// REFRESH TOKEN
	refreshExp := time.Now().Add(time.Hour * 24 * time.Duration(s.cfg.RefreshExpireDays))

	refreshClaims := jwt.RegisteredClaims{
		Subject:   user.Username,
		ExpiresAt: jwt.NewNumericDate(refreshExp),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenStr, err := refreshToken.SignedString([]byte(s.cfg.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return &entity.TokenResponse{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
		TokenType:    "Bearer",
		ExpiresIn:    int64(s.cfg.AccessExpireMinutes * 60),
		ExpiresDate:  refreshExp,
	}, nil
}
