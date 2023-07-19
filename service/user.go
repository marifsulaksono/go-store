package service

import (
	"errors"
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"

	"gorm.io/gorm"
)

type UserService struct {
	Repo repo.UserRepository
}

func NewUserService(r repo.UserRepository) *UserService {
	return &UserService{
		Repo: r,
	}
}

func (u UserService) GetUserByUsername(username string) (entity.UserResponse, error) {
	user, err := u.Repo.GetUserByUsername(username)
	return user, err
}

func (u UserService) CreateUser(user *entity.User) error {
	userCheck, err := u.Repo.GetUserByUsername(user.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = u.Repo.CreateUser(user)
		return err
	} else if user.Username == userCheck.Username {
		return helper.ErrUserExist
	}
	return err
}
