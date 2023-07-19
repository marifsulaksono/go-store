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
