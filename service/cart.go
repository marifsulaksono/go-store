package service

import (
	"context"
	"errors"
	"gostore/entity"
	"gostore/helper"
	"gostore/middleware"
	"gostore/repo"

	"gorm.io/gorm"
)

type cartService struct {
	Repo        repo.CartRepository
	ProductRepo repo.ProductRepository
}

type CartService interface {
	GetCart(ctx context.Context) ([]entity.Cart, error)
	CreateCart(ctx context.Context, cart *entity.Cart) error
	UpdateCart(ctx context.Context, id int, cart *entity.Cart) error
	DeleteCart(ctx context.Context, id int) error
}

func NewCartService(r repo.CartRepository, p repo.ProductRepository) CartService {
	return &cartService{
		Repo:        r,
		ProductRepo: p,
	}
}

func (c *cartService) GetCart(ctx context.Context) ([]entity.Cart, error) {
	return c.Repo.GetCart(ctx)
}

func (c *cartService) CreateCart(ctx context.Context, cart *entity.Cart) error {
	userId := ctx.Value(middleware.GOSTORE_USERID).(int)
	product, err := c.ProductRepo.GetProductById(ctx, cart.ProductId)
	if err != nil {
		return helper.ErrProductNotFound
	}

	// check stock & valid product
	if product.Stock < cart.Qty {
		return helper.ErrStockNotEnough
	} else if product.Store.UserId == userId {
		return errors.New("cannot add your product's store to your cart")
	}

	checkCart, err := c.Repo.GetCartId(ctx, cart.ProductId, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// create the cart if not productid on userid cart
			cart.UserId = userId
			err = c.Repo.CreateCart(ctx, cart)
			return err
		}
		return err
	}

	// supplement total qty if the product added is same
	if cart.ProductId == checkCart.ProductId {
		cart.Qty += checkCart.Qty
	}

	return c.Repo.UpdateCart(ctx, checkCart.Id, cart)
}

func (c *cartService) UpdateCart(ctx context.Context, id int, cart *entity.Cart) error {
	checkCart, err := c.Repo.GetCartById(ctx, id)
	if err != nil {
		return err
	}

	checkProduct, err := c.ProductRepo.GetProductById(ctx, checkCart.ProductId)
	if err != nil {
		return err
	} else if checkProduct.Stock < cart.Qty {
		return helper.ErrStockNotEnough
	}

	return c.Repo.UpdateCart(ctx, id, cart)
}

func (c *cartService) DeleteCart(ctx context.Context, id int) error {
	_, err := c.Repo.GetCartById(ctx, id)
	if err != nil {
		return err
	}

	return c.Repo.DeleteCart(ctx, id)
}
