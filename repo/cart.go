package repo

import (
	"gostore/entity"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		DB: db,
	}
}

func (c *CartRepository) GetCart(userId int) ([]entity.Cart, error) {
	var result []entity.Cart
	err := c.DB.Where("user_id = ?", userId).Preload("Product.Store").Find(&result).Error
	return result, err
}

func (c *CartRepository) GetCartById(id int) (entity.Cart, error) {
	var result entity.Cart
	err := c.DB.Where("id = ?", id).Find(&result).Error
	return result, err
}

func (c *CartRepository) CreateCart(cart *entity.Cart) error {
	err := c.DB.Create(cart).Error
	return err
}

func (c *CartRepository) UpdateCart(id int, model, cart *entity.Cart) error {
	err := c.DB.Model(model).Where("id = ?", id).Updates(cart).Error
	return err
}

func (c *CartRepository) DeleteCart(id int) error {
	err := c.DB.Where("id = ?", id).Delete(&entity.Cart{}).Error
	return err
}
