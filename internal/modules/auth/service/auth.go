package service

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/waldin-exe/golang-bootstrap/config"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/auth/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/auth/entity"
	appErrors "github.com/waldin-exe/golang-bootstrap/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo         contract.AuthRepository
	tokenService contract.TokenService
	cfg          config.Config
}

func NewAuthService(repo contract.AuthRepository, tokenService contract.TokenService, cfg config.Config) contract.AuthService {
	return &authService{repo: repo, tokenService: tokenService, cfg: cfg}
}

func (s *authService) Login(ctx context.Context, req entity.LoginRequest) (*entity.TokenResponse, error) {
	user, err := s.repo.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, appErrors.NewNotFoundError("User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, appErrors.NewBadRequestError("Invalid password")
	}

	claimsJWT := entity.AuthUserClaims{
		Username:  user.Username,
		Role:      user.Level,
		UserID:    user.Id,
		PegawaiID: user.PegawaiID,
	}

	// Delete old refresh tokens
	_ = s.repo.DeleteRefreshTokensByEmail(ctx, user.Username)

	token, err := s.tokenService.GenerateToken(claimsJWT)
	if err != nil {
		return nil, appErrors.NewInternalError("Failed to generate token")
	}

	if err := s.repo.SaveRefreshToken(ctx, user.Username, token.RefreshToken, token.ExpiresDate); err != nil {
		return nil, appErrors.NewInternalError("Failed to save refresh token")
	}

	return token, nil
}

func (s *authService) RefreshToken(ctx context.Context, req entity.RefreshTokenRequest) (*entity.TokenResponse, error) {
	refreshSecret := s.cfg.JWT.RefreshSecret

	claims := &jwt.RegisteredClaims{}
	jwtToken, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(refreshSecret), nil
	})

	if err != nil || !jwtToken.Valid {
		return nil, appErrors.NewUnauthorizedError("Invalid or expired refresh token")
	}

	email := claims.Subject
	valid, err := s.repo.CheckRefreshTokenValid(ctx, email, req.RefreshToken)
	if err != nil || !valid {
		return nil, appErrors.NewUnauthorizedError("Refresh token not recognized")
	}

	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, appErrors.NewInternalError("User not found for this token")
	}

	claimsJWT := entity.AuthUserClaims{
		UserID:    user.Id,
		Username:  user.Username,
		Role:      user.Level,
		PegawaiID: user.PegawaiID,
	}

	tokens, err := s.tokenService.GenerateToken(claimsJWT)
	if err != nil {
		return nil, appErrors.NewInternalError("Failed to generate token")
	}

	if err := s.repo.UpdateRefreshToken(ctx, email, tokens.RefreshToken, tokens.ExpiresDate); err != nil {
		return nil, appErrors.NewInternalError("Failed to update refresh token")
	}

	return tokens, nil
}
