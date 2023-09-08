package repo

import (
	"context"
	"errors"
	"gostore/entity"
	userError "gostore/helper/domain/errorModel"
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	GetUser(ctx context.Context, id int, username string) (entity.UserResponse, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, id int, user *entity.User) error
	ChangePasswordUser(ctx context.Context, id int, password string) error
	DeleteUser(ctx context.Context, id int) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) GetUser(ctx context.Context, id int, username string) (entity.UserResponse, error) {
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

	db.Where("delete_at is null")
	err := db.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.UserResponse{}, userError.ErrUserNotFound
		}
		return entity.UserResponse{}, err
	}
	return user, nil
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	user.Role = "Buyer"
	return u.DB.Create(user).Error
}

func (u *userRepository) UpdateUser(ctx context.Context, id int, user *entity.User) error {
	user.UpdateAt = time.Now()
	return u.DB.Model(entity.User{}).Where("id = ?", id).Updates(&user).Error
}

func (u *userRepository) ChangePasswordUser(ctx context.Context, id int, password string) error {
	return u.DB.Model(entity.User{}).Where("id = ?", id).Update("password", password).Error
}

func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	return u.DB.Where("id = ?", id).Delete(&entity.User{}).Error
}
