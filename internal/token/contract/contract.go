package contract

import "github.com/waldin-exe/golang-bootstrap/internal/modules/auth/entity"

type TokenService interface {
	GenerateToken(user entity.AuthUserClaims) (*entity.TokenResponse, error)
}
