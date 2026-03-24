package contract

import (
	"context"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/user/entity"
)

type UserRepository interface {
	GetUsers(ctx context.Context, filter entity.GetUserRequest) ([]entity.UserResponse, int64, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int, updatedBy string) error
	FindByID(ctx context.Context, id int) (*entity.User, error)
}

type UserService interface {
	GetUsers(ctx context.Context, req entity.GetUserRequest) ([]entity.UserResponse, int64, error)
	CreateUser(ctx context.Context, req entity.PostUserRequest) error
	UpdateUser(ctx context.Context, req entity.PutUserRequest) error
	DeleteUser(ctx context.Context, req entity.DeleteUserRequest) error
	UbahPassword(ctx context.Context, req entity.UbahPasswordInput) error
}
