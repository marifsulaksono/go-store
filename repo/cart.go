package repo

import (
	"context"
	"errors"
	"gostore/entity"
	"gostore/utils/helper"
	cartError "gostore/utils/helper/domain/errorModel"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

type CartRepository interface {
	GetCart(ctx context.Context) ([]entity.Cart, error)
	GetCartById(ctx context.Context, id int) (entity.Cart, error)
	GetCartProductId(ctx context.Context, productId, userId int) (entity.Cart, error)
	CreateCart(ctx context.Context, cart *entity.Cart) error
	UpdateCart(ctx context.Context, id int, cart *entity.Cart) error
	DeleteCart(ctx context.Context, id int) error
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (c *cartRepository) GetCart(ctx context.Context) ([]entity.Cart, error) {
	userId := ctx.Value(helper.GOSTORE_USERID).(int)
	var result []entity.Cart
	err := c.DB.Where("user_id = ?", userId).Preload("Product.Store").Find(&result).Error
	return result, err
}

func (c *cartRepository) GetCartById(ctx context.Context, id int) (entity.Cart, error) {
	userId := ctx.Value(helper.GOSTORE_USERID)
	var result entity.Cart
	err := c.DB.Where("id = ? and user_id = ?", id, userId).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Cart{}, cartError.ErrCartNotFound
		}
		return entity.Cart{}, err
	}
	return result, nil
}

// get cart id by product id and user id
func (c *cartRepository) GetCartProductId(ctx context.Context, productId, userId int) (entity.Cart, error) {
	var result entity.Cart
	err := c.DB.Where("product_id = ? and user_id = ?", productId, userId).First(&result).Error
	return result, err
}

func (c *cartRepository) CreateCart(ctx context.Context, cart *entity.Cart) error {
	return c.DB.Create(cart).Error
}

func (c *cartRepository) UpdateCart(ctx context.Context, id int, cart *entity.Cart) error {
	return c.DB.Model(&entity.Cart{}).Where("id = ?", id).Update("qty", *cart.Qty).Error
}

func (c *cartRepository) DeleteCart(ctx context.Context, id int) error {
	return c.DB.Where("id = ?", id).Delete(&entity.Cart{}).Error
}
