package service

import (
	"context"
	"errors"
	"gostore/entity"
	cartError "gostore/helper/domain/errorModel"
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
	var (
		detailError = make(map[string]any)
		userId      = ctx.Value(middleware.GOSTORE_USERID).(int)
	)

	// input validation
	if cart.ProductId == nil {
		detailError["product_id"] = "this field is missing input"
	} else if *cart.ProductId < 0 {
		detailError["product_id"] = "this field must not negative"
	}

	if cart.Qty == nil {
		detailError["qty"] = "this field is missing input"
	} else if *cart.Qty < 0 {
		detailError["qty"] = "this field must not negative"
	}

	if len(detailError) > 0 {
		return cartError.ErrCartInput.AttachDetail(detailError)
	}

	product, err := c.ProductRepo.GetProductById(ctx, *cart.ProductId)
	if err != nil {
		return err
	}

	// check stock & valid product
	if *product.Stock < *cart.Qty {
		return cartError.ErrStockProductNotEnough
	} else if product.Store.UserId == userId {
		return cartError.ErrCantAddToCart
	}

	// check by by product id and user id to create new cart or update the exist cart
	checkCart, err := c.Repo.GetCartProductId(ctx, *cart.ProductId, userId)
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
	if *cart.ProductId == *checkCart.ProductId {
		*cart.Qty += *checkCart.Qty
	}

	return c.Repo.UpdateCart(ctx, checkCart.Id, cart)
}

func (c *cartService) UpdateCart(ctx context.Context, id int, cart *entity.Cart) error {
	var (
		detailError = make(map[string]any)
	)

	// input validation
	if cart.ProductId == nil {
		detailError["product_id"] = "this field is missing input"
	} else if *cart.ProductId < 0 {
		detailError["product_id"] = "this field must not negative"
	}

	if cart.Qty == nil {
		detailError["qty"] = "this field is missing input"
	} else if *cart.Qty < 0 {
		detailError["qty"] = "this field must not negative"
	}

	if len(detailError) > 0 {
		return cartError.ErrCartInput.AttachDetail(detailError)
	}

	checkCart, err := c.Repo.GetCartById(ctx, id)
	if err != nil {
		return err
	}

	checkProduct, err := c.ProductRepo.GetProductById(ctx, *checkCart.ProductId)
	if err != nil {
		return err
	} else if *checkProduct.Stock < *cart.Qty {
		return cartError.ErrStockProductNotEnough
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
