package repository

import (
	"context"
	"errors"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/user/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/user/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) contract.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers(ctx context.Context, filter entity.GetUserRequest) ([]entity.UserResponse, int64, error) {
	var users []entity.UserResponse
	var totalData int64

	query := r.db.WithContext(ctx).Table("users").
		Select("users.id, pegawai.nama as nama_pegawai, users.username, users.level").
		Joins("INNER JOIN pegawais pegawai ON pegawai.id = users.pegawai_id").
		Where("users.deleted_at IS NULL")

	if filter.ID != 0 {
		query = query.Where("users.id = ?", filter.ID)
	}
	if filter.NamaPegawai != "" {
		query = query.Where("pegawai.nama ILIKE ?", "%"+filter.NamaPegawai+"%")
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

	err = query.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, totalData, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil, nil when not found is standard
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) DeleteUser(ctx context.Context, id int, updatedBy string) error {
	result := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"updated_by": updatedBy,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return r.db.WithContext(ctx).Delete(&entity.User{}, id).Error
}
