package repo

import (
	"context"
	"errors"
	"gostore/entity"
	"gostore/utils/helper"
	saError "gostore/utils/helper/domain/errorModel"

	"gorm.io/gorm"
)

type shippingAddressRepo struct {
	DB *gorm.DB
}

type ShippingAddressRepo interface {
	GetShippingAddressById(ctx context.Context, id int) (entity.ShippingAddress, error)
	GetShippingAddressByUserId(ctx context.Context) ([]entity.ShippingAddress, error)
	InsertShippingAddress(ctx context.Context, sa *entity.ShippingAddress) error
	UpdateShippingAddress(ctx context.Context, id int, model, sa *entity.ShippingAddress) error
	DeleteShippingAddress(ctx context.Context, id int) error
}

func NewShippingAddressRepo(db *gorm.DB) ShippingAddressRepo {
	return &shippingAddressRepo{
		DB: db,
	}
}

func (u *shippingAddressRepo) GetShippingAddressById(ctx context.Context, id int) (entity.ShippingAddress, error) {
	var SA entity.ShippingAddress
	err := u.DB.Where("id = ?", id).First(&SA).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.ShippingAddress{}, saError.ErrAddressNotFound
		}
		return entity.ShippingAddress{}, err
	}
	return SA, err
}

func (u *shippingAddressRepo) GetShippingAddressByUserId(ctx context.Context) ([]entity.ShippingAddress, error) {
	userId := ctx.Value(helper.GOSTORE_USERID).(int)
	var SA []entity.ShippingAddress
	err := u.DB.Where("user_id = ?", userId).Find(&SA).Error
	return SA, err
}

func (u *shippingAddressRepo) InsertShippingAddress(ctx context.Context, sa *entity.ShippingAddress) error {
	sa.UserId = ctx.Value(helper.GOSTORE_USERID).(int)
	return u.DB.Create(sa).Error
}

func (u *shippingAddressRepo) UpdateShippingAddress(ctx context.Context, id int, model, sa *entity.ShippingAddress) error {
	return u.DB.Model(model).Where("id = ?", id).Updates(sa).Error
}

func (u *shippingAddressRepo) DeleteShippingAddress(ctx context.Context, id int) error {
	return u.DB.Where("id = ?", id).Delete(&entity.ShippingAddress{}).Error
}
