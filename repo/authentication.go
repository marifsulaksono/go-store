package repo

import (
	"context"
	"gostore/entity"
	authError "gostore/utils/helper/domain/errorModel"

	"gorm.io/gorm"
)

type authentiocationRepo struct {
	DB *gorm.DB
}

type AuthentiocationRepo interface {
	CreateSession(ctx context.Context, s *entity.Session) error
	Logout(ctx context.Context, token string) error
	ValidateRefreshToken(ctx context.Context, token string) error
}

func NewAuthentiocationRepo(db *gorm.DB) AuthentiocationRepo {
	return &authentiocationRepo{DB: db}
}

func (a *authentiocationRepo) CreateSession(ctx context.Context, s *entity.Session) error {
	tx := a.DB.Create(s)
	return tx.Error
}

func (a *authentiocationRepo) Logout(ctx context.Context, token string) error {
	tx := a.DB.Where("token = ?", token).Delete(&entity.Session{})
	if tx.RowsAffected < 1 {
		return authError.ErrInvalidToken
	}

	return tx.Error
}

func (a *authentiocationRepo) ValidateRefreshToken(ctx context.Context, token string) error {
	var result entity.Session
	tx := a.DB.Where("token = ?", token).First(&result)
	return tx.Error
}
