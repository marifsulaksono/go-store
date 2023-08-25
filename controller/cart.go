package controller

import (
	"encoding/json"
	"gostore/entity"
	"gostore/helper"
	"gostore/service"
	"net/http"
)

type CartController struct {
	Service service.CartService
}

func NewCartController(s service.CartService) *CartController {
	return &CartController{
		Service: s,
	}
}

func (c *CartController) GetCartByUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := c.Service.GetCart(ctx)
	if err != nil {
		helper.BuildError(w, err)
		return
	}

	if result == nil || len(result) < 1 {
		helper.BuildResponseSuccess(w, result, nil, "No results found")
		return
	}

	helper.BuildResponseSuccess(w, result, nil, "")
}

func (c *CartController) CreateCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var cart entity.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		helper.BuildError(w, err)
		return
	}
	defer r.Body.Close()

	err = c.Service.CreateCart(ctx, &cart)
	if err != nil {
		helper.BuildError(w, err)
		return
	}

	helper.BuildResponseSuccess(w, nil, nil, "Success add new cart")
}

func (c *CartController) UpdateCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		var cart entity.Cart
		err := json.NewDecoder(r.Body).Decode(&cart)
		if err != nil {
			helper.BuildError(w, err)
			return
		}

		err = c.Service.UpdateCart(ctx, id, &cart)
		if err != nil {
			helper.BuildError(w, err)
			return
		}

		helper.BuildResponseSuccess(w, nil, nil, "Success update cart")
	}
}

func (c *CartController) DeleteCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		err := c.Service.DeleteCart(ctx, id)
		if err != nil {
			helper.BuildError(w, err)
			return
		}

		helper.BuildResponseSuccess(w, nil, nil, "Success delete cart")
	}
}
