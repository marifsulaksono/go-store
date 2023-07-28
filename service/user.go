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

func (u *UserService) GetUserByUsername(username string) (entity.UserResponse, error) {
	user, err := u.Repo.GetUserByUsername(username)
	return user, err
}

func (u *UserService) CreateUser(user *entity.User) error {
	userCheck, err := u.Repo.GetUserByUsername(user.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = u.Repo.CreateUser(user)
		return err
	} else if user.Username == userCheck.Username {
		return helper.ErrUserExist
	}
	return err
}

func (u *UserService) GetShippingAddressByUserId(userId int) ([]entity.ShippingAddress, error) {
	SA, err := u.Repo.GetShippingAddressByUserId(userId)
	return SA, err
}

func (u *UserService) InsertShippingAddress(sa *entity.ShippingAddress) error {
	err := u.Repo.InsertShippingAddress(sa)
	return err
}

func (u *UserService) UpdateShippingAddress(userId, id int, sa *entity.ShippingAddress) (entity.ShippingAddress, error) {
	existSA, err := u.Repo.GetShippingAddressById(id)
	if err != nil {
		return entity.ShippingAddress{}, err
	}

	if existSA.UserId != userId {
		return entity.ShippingAddress{}, helper.ErrInvalidUser
	}
	err = u.Repo.UpdateShippingAddress(id, &existSA, sa)
	if err != nil {
		return entity.ShippingAddress{}, err
	} else if existSA, err = u.Repo.GetShippingAddressById(id); err != nil {
		return entity.ShippingAddress{}, err
	}
	return existSA, err
}

func (u *UserService) DeleteShippingAddress(userId, id int) error {
	existSA, err := u.Repo.GetShippingAddressById(id)
	if err != nil {
		return err
	}

	if existSA.UserId != userId {
		return helper.ErrInvalidUser
	}
	err = u.Repo.DeleteShippingAddress(id)
	return err
}
