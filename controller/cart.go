package controller

import (
	"encoding/json"
	"fmt"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if len(result) == 0 {
		helper.ResponseWrite(w, result, fmt.Sprintln("Data not found"))
		return
	}

	helper.ResponseWrite(w, result, "Success get all carts")
}

func (c *CartController) CreateCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var cart entity.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = c.Service.CreateCart(ctx, &cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, cart, "Success create new cart")
}

func (c *CartController) UpdateCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		var cart entity.Cart
		err := json.NewDecoder(r.Body).Decode(&cart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = c.Service.UpdateCart(ctx, id, &cart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, id, fmt.Sprintf("Success update cart %d on user", id))
	}
}

func (c *CartController) DeleteCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		err := c.Service.DeleteCart(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, id, fmt.Sprintf("Success delete cart %d on user", id))
	}
}
