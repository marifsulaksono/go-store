package service

import (
	"context"
	"gostore/entity"
	"gostore/helper"
	"gostore/middleware"
	"gostore/repo"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	Repo repo.UserRepository
}

type UserService interface {
	GetUser(id int, username string) (entity.UserResponse, error)
	CreateUser(user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	ChangePasswordUser(ctx context.Context, userChange entity.UserChangePassword) error
	DeleteUser(ctx context.Context) error
	GetShippingAddressByUserId(ctx context.Context) ([]entity.ShippingAddress, error)
	InsertShippingAddress(ctx context.Context, sa *entity.ShippingAddress) error
	UpdateShippingAddress(ctx context.Context, id int, sa *entity.ShippingAddress) error
	DeleteShippingAddress(ctx context.Context, id int) error
}

func NewUserService(r repo.UserRepository) UserService {
	return &userService{
		Repo: r,
	}
}

func (u *userService) GetUser(id int, username string) (entity.UserResponse, error) {
	return u.Repo.GetUser(id, username)
}

func (u *userService) CreateUser(user *entity.User) error {
	// generate password entry to be encrypt
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPwd)
	user.CreateAt = time.Now()

	err := u.Repo.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			return helper.ErrUserExist
		}
		return err
	}
	return nil
}

func (u *userService) UpdateUser(ctx context.Context, user *entity.User) error {
	return u.Repo.UpdateUser(ctx, user)
}

func (u *userService) ChangePasswordUser(ctx context.Context, userChange entity.UserChangePassword) error {
	userId := ctx.Value(middleware.GOSTORE_USERID).(int)
	checkUser, err := u.Repo.GetUser(userId, "")
	if err != nil {
		return err
	}

	// validation old password
	if err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(userChange.OldPassword)); err != nil {
		return helper.ErrWrongOldPassword
	}

	// generate hash of new password
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(userChange.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userChange.NewPassword = string(hashPwd)
	return u.Repo.ChangePasswordUser(ctx, userId, userChange.NewPassword)
}

func (u *userService) DeleteUser(ctx context.Context) error {
	id := ctx.Value(middleware.GOSTORE_USERID).(int)
	_, err := u.Repo.GetUser(id, "")
	if err != nil {
		return err
	}

	return u.Repo.DeleteUser(ctx)
}

func (u *userService) GetShippingAddressByUserId(ctx context.Context) ([]entity.ShippingAddress, error) {
	return u.Repo.GetShippingAddressByUserId(ctx)
}

func (u *userService) InsertShippingAddress(ctx context.Context, sa *entity.ShippingAddress) error {
	return u.Repo.InsertShippingAddress(ctx, sa)
}

func (u *userService) UpdateShippingAddress(ctx context.Context, id int, sa *entity.ShippingAddress) error {
	userId := ctx.Value(middleware.GOSTORE_USERID).(int)
	existSA, err := u.Repo.GetShippingAddressById(ctx, id)
	if err != nil {
		return err
	}

	// validation user updater is right user data
	if existSA.UserId != userId {
		return helper.ErrInvalidUser
	}
	err = u.Repo.UpdateShippingAddress(ctx, id, &existSA, sa)
	return err
}

func (u *userService) DeleteShippingAddress(ctx context.Context, id int) error {
	userId := ctx.Value(middleware.GOSTORE_USERID).(int)
	existSA, err := u.Repo.GetShippingAddressById(ctx, id)
	if err != nil {
		return err
	}

	if existSA.UserId != userId {
		return helper.ErrInvalidUser
	}

	return u.Repo.DeleteShippingAddress(ctx, id)
}
