package repo

import (
	"gostore/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) GetUserByUsername(username string) (entity.UserResponse, error) {
	var user entity.UserResponse
	err := u.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (u *UserRepository) CreateUser(user *entity.User) error {
	err := u.DB.Create(user).Error
	return err
}

func (u *UserRepository) GetShippingAddressById(id int) (entity.ShippingAddress, error) {
	var SA entity.ShippingAddress
	err := u.DB.Where("id = ?", id).First(&SA).Error
	return SA, err
}

func (u *UserRepository) GetShippingAddressByUserId(userId int) ([]entity.ShippingAddress, error) {
	var SA []entity.ShippingAddress
	err := u.DB.Where("user_id = ?", userId).Find(&SA).Error
	return SA, err
}

func (u *UserRepository) InsertShippingAddress(sa *entity.ShippingAddress) error {
	err := u.DB.Create(sa).Error
	return err
}

func (u *UserRepository) UpdateShippingAddress(id int, model, sa *entity.ShippingAddress) error {
	err := u.DB.Model(model).Where("id = ?", id).Updates(sa).Error
	return err
}

func (u *UserRepository) DeleteShippingAddress(id int) error {
	err := u.DB.Where("id = ?", id).Delete(&entity.ShippingAddress{}).Error
	return err
}
