package repo

import (
	"context"
	"gostore/entity"
	"gostore/middleware"
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	GetUser(id int, username string) (entity.UserResponse, error)
	CreateUser(user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	ChangePasswordUser(ctx context.Context, id int, password string) error
	DeleteUser(ctx context.Context) error
	GetShippingAddressById(ctx context.Context, id int) (entity.ShippingAddress, error)
	GetShippingAddressByUserId(ctx context.Context) ([]entity.ShippingAddress, error)
	InsertShippingAddress(ctx context.Context, sa *entity.ShippingAddress) error
	UpdateShippingAddress(ctx context.Context, id int, model, sa *entity.ShippingAddress) error
	DeleteShippingAddress(ctx context.Context, id int) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) GetUser(id int, username string) (entity.UserResponse, error) {
	var (
		user entity.UserResponse
		db   = u.DB
	)

	if id != 0 && id > 0 {
		db = db.Where("id = ?", id)
	}

	if username != "" {
		db = db.Where("username = ?", username)
	}

	err := db.First(&user).Error
	return user, err
}

func (u *userRepository) CreateUser(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	id := ctx.Value(middleware.GOSTORE_USERID).(int)
	user.UpdateAt = time.Now()
	return u.DB.Model(entity.User{}).Where("id = ?", id).Updates(&user).Error
}

func (u *userRepository) ChangePasswordUser(ctx context.Context, id int, password string) error {
	return u.DB.Model(entity.User{}).Where("id = ?", id).Update("password", password).Error
}

func (u *userRepository) DeleteUser(ctx context.Context) error {
	id := ctx.Value(middleware.GOSTORE_USERID).(int)
	return u.DB.Where("id = ?", id).Delete(entity.User{}).Error
}

func (u *userRepository) GetShippingAddressById(ctx context.Context, id int) (entity.ShippingAddress, error) {
	var SA entity.ShippingAddress
	err := u.DB.Where("id = ?", id).First(&SA).Error
	return SA, err
}

func (u *userRepository) GetShippingAddressByUserId(ctx context.Context) ([]entity.ShippingAddress, error) {
	userId := ctx.Value(middleware.GOSTORE_USERID).(int)
	var SA []entity.ShippingAddress
	err := u.DB.Where("user_id = ?", userId).Find(&SA).Error
	return SA, err
}

func (u *userRepository) InsertShippingAddress(ctx context.Context, sa *entity.ShippingAddress) error {
	sa.UserId = ctx.Value(middleware.GOSTORE_USERID).(int)
	return u.DB.Create(sa).Error
}

func (u *userRepository) UpdateShippingAddress(ctx context.Context, id int, model, sa *entity.ShippingAddress) error {
	return u.DB.Model(model).Where("id = ?", id).Updates(sa).Error
}

func (u *userRepository) DeleteShippingAddress(ctx context.Context, id int) error {
	return u.DB.Where("id = ?", id).Delete(&entity.ShippingAddress{}).Error
}
