package service

import (
	"context"

	"github.com/waldin-exe/golang-bootstrap/internal/modules/user/contract"
	"github.com/waldin-exe/golang-bootstrap/internal/modules/user/entity"
	appErrors "github.com/waldin-exe/golang-bootstrap/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo contract.UserRepository
}

func NewUserService(repo contract.UserRepository) contract.UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUsers(ctx context.Context, req entity.GetUserRequest) ([]entity.UserResponse, int64, error) {
	users, count, err := s.repo.GetUsers(ctx, req)
	if err != nil {
		return nil, 0, appErrors.NewInternalError(err.Error())
	}
	if users == nil {
		users = []entity.UserResponse{}
	}
	return users, count, nil
}

func (s *userService) CreateUser(ctx context.Context, req entity.PostUserRequest) error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return appErrors.NewInternalError("Failed to hash password")
	}

	user := &entity.User{
		Username:  req.Username,
		Password:  string(hashedPwd),
		Level:     req.Level,
		PegawaiID: req.PegawaiId,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.CreatedBy,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	return nil
}

func (s *userService) UpdateUser(ctx context.Context, req entity.PutUserRequest) error {
	user, err := s.repo.FindByID(ctx, req.Id)
	if err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	if user == nil {
		return appErrors.NewNotFoundError("User not found")
	}

	user.Username = req.Username
	user.Level = req.Level
	user.UpdatedBy = req.UpdatedBy

	if req.UpdatedBy == "1" {
		// Admin changing someone else's password
		if req.NewPassword != "" {
			hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
			user.Password = string(hashedPwd)
		}
	} else if req.OldPassword != "" {
		// User changing their own password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
			return appErrors.NewBadRequestError("Old password is incorrect")
		}
		if req.NewPassword != "" {
			hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
			user.Password = string(hashedPwd)
		}
	}

	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	return nil
}

func (s *userService) DeleteUser(ctx context.Context, req entity.DeleteUserRequest) error {
	user, err := s.repo.FindByID(ctx, req.Id)
	if err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	if user == nil {
		return appErrors.NewNotFoundError("User not found")
	}
	if err := s.repo.DeleteUser(ctx, req.Id, req.UpdatedBy); err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	return nil
}

func (s *userService) UbahPassword(ctx context.Context, req entity.UbahPasswordInput) error {
	user, err := s.repo.FindByID(ctx, req.UserID)
	if err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	if user == nil {
		return appErrors.NewNotFoundError("User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.CurrentPassword)); err != nil {
		return appErrors.NewBadRequestError("Current password is incorrect")
	}

	newHashed, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return appErrors.NewInternalError("Failed to hash new password")
	}

	user.Password = string(newHashed)

	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return appErrors.NewInternalError(err.Error())
	}
	return nil
}
