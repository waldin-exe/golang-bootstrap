package repository

import (
	"context"
	"time"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/auth/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/auth/entity"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) contract.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindUserByUsername(ctx context.Context, username string) (*contract.AuthUser, error) {
	var user contract.AuthUser
	query := r.db.WithContext(ctx).Table("users as u").
		Select("u.id, u.password, u.username, p.nama, p.tgl_lahir, p.alamat, p.no_telepon, u.level, u.pegawai_id").
		Joins("INNER JOIN pegawais as p ON p.id = u.pegawai_id").
		Where("u.username = ? AND p.deleted_at IS NULL AND u.deleted_at IS NULL", username)

	if err := query.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) FindUserByEmail(ctx context.Context, email string) (*contract.AuthUser, error) {
	var user contract.AuthUser
	query := r.db.WithContext(ctx).Table("users as u").
		Select("u.id, u.username, p.nama, p.tgl_lahir, p.alamat, p.no_telepon, u.level, u.pegawai_id").
		Joins("INNER JOIN pegawais as p ON p.id = u.pegawai_id").
		Where("u.username = ?", email) // In the old code "FindUserByEmail" queries username with email argument

	if err := query.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) SaveRefreshToken(ctx context.Context, email, token string, expiredAt time.Time) error {
	rt := &entity.RefreshToken{
		UserEmail: email,
		Token:     token,
		ExpiredAt: expiredAt,
	}
	return r.db.WithContext(ctx).Create(rt).Error
}

func (r *authRepository) DeleteRefreshTokensByEmail(ctx context.Context, email string) error {
	return r.db.WithContext(ctx).Where("user_email = ?", email).Delete(&entity.RefreshToken{}).Error
}

func (r *authRepository) CheckRefreshTokenValid(ctx context.Context, email, token string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.RefreshToken{}).
		Where("user_email = ? AND token = ? AND expired_at > ?", email, token, time.Now()).
		Count(&count).Error
	return count > 0, err
}

func (r *authRepository) UpdateRefreshToken(ctx context.Context, email, token string, expiredAt time.Time) error {
	return r.db.WithContext(ctx).Model(&entity.RefreshToken{}).
		Where("user_email = ?", email).
		Updates(map[string]interface{}{
			"token":      token,
			"expired_at": expiredAt,
		}).Error
}
