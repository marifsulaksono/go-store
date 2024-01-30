package controller

import (
	"encoding/json"
	"gostore/entity"
	"gostore/service"
	"gostore/utils/helper"
	"gostore/utils/response"
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
	var (
		ctx     = r.Context()
		message string
	)

	result, err := c.Service.GetCart(ctx)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if result == nil || len(result) < 1 {
		message = "No results found"
	}

	response.BuildSuccesResponse(w, result, nil, message)
}

func (c *CartController) CreateCart(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		cart entity.Cart
	)

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	err = c.Service.CreateCart(ctx, &cart)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success add new cart")
}

func (c *CartController) UpdateCart(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		cart entity.Cart
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	err = c.Service.UpdateCart(ctx, id, &cart)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success update cart")
}

func (c *CartController) DeleteCart(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := c.Service.DeleteCart(ctx, id); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success delete cart")
}
