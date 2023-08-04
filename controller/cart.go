package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/middleware"
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
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)
	result, err := c.Service.GetCart(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, fmt.Sprintf("Success get all carts on user %d", userId))
}

func (c *CartController) CreateCart(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)
	var cart entity.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	cart.UserId = userId
	err = c.Service.CreateCart(userId, &cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, cart, fmt.Sprintf("Success create new cart on user %d", userId))
}

func (c *CartController) UpdateCart(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)
	if id, s := helper.IdVarsMux(w, r); s {
		var cart entity.Cart
		err := json.NewDecoder(r.Body).Decode(&cart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = c.Service.UpdateCart(userId, id, &cart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, id, fmt.Sprintf("Success update cart %d on user %d", id, userId))
	}
}

func (c *CartController) DeleteCart(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)
	if id, s := helper.IdVarsMux(w, r); s {
		err := c.Service.DeleteCart(userId, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, id, fmt.Sprintf("Success delete cart %d on user %d", id, userId))
	}
}
