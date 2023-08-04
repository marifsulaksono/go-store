package service

import (
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"
)

type CartService struct {
	Repo        repo.CartRepository
	ProductRepo repo.ProductRepository
}

func NewCartService(r repo.CartRepository, p repo.ProductRepository) *CartService {
	return &CartService{
		Repo:        r,
		ProductRepo: p,
	}
}

func (c *CartService) GetCart(userId int) ([]entity.Cart, error) {
	result, err := c.Repo.GetCart(userId)
	return result, err
}

func (c *CartService) CreateCart(userId int, cart *entity.Cart) error {
	product, err := c.ProductRepo.GetProductById(cart.ProductId)
	if err != nil {
		return err
	}

	// check stock product
	if product.Stock < cart.Qty {
		return helper.ErrStockNotEnough
	}

	err = c.Repo.CreateCart(cart)
	return err
}

func (c *CartService) UpdateCart(userId, id int, cart *entity.Cart) error {
	checkCart, err := c.Repo.GetCartById(id)
	if err != nil {
		return err
	} else if checkCart.UserId != userId {
		return helper.ErrAccDeny
	}

	checkProduct, err := c.ProductRepo.GetProductById(checkCart.ProductId)
	if err != nil {
		return err
	} else if checkProduct.Stock < cart.Qty {
		return helper.ErrStockNotEnough
	}

	err = c.Repo.UpdateCart(id, &checkCart, cart)
	return err
}

func (c *CartService) DeleteCart(userId, id int) error {
	checkCart, err := c.Repo.GetCartById(id)
	if err != nil {
		return err
	} else if checkCart.UserId != userId {
		return helper.ErrAccDeny
	}

	err = c.Repo.DeleteCart(id)
	return err
}
