package controller

import (
	"encoding/json"
	"gostore/entity"
	"gostore/helper"
	"gostore/helper/response"
	"gostore/service"
	"net/http"
)

type ShippingAddressController struct {
	Service service.ShippingAddressService
}

func NewShippingAddressController(service service.ShippingAddressService) *ShippingAddressController {
	return &ShippingAddressController{
		Service: service,
	}
}

func (u *ShippingAddressController) GetShippingAddressByUserId(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		message string
	)

	result, err := u.Service.GetShippingAddressByUserId(ctx)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if result == nil || len(result) < 1 {
		message = "No results found"
	}

	response.BuildSuccesResponse(w, result, nil, message)
}

func (u *ShippingAddressController) GetShippingAddressById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	result, err := u.Service.GetShippingAddressById(ctx, id)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, result, nil, "")
}

func (u *ShippingAddressController) InsertShippingAddress(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		SA  entity.ShippingAddress
	)

	err := json.NewDecoder(r.Body).Decode(&SA)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	if err := u.Service.InsertShippingAddress(ctx, &SA); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success create new shipping address")
}

func (u *ShippingAddressController) UpdateShippingAddress(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		SA  entity.ShippingAddress
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&SA)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	if err = u.Service.UpdateShippingAddress(ctx, id, &SA); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success update shipping address")
}

func (u *ShippingAddressController) DeleteShippingAddress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := u.Service.DeleteShippingAddress(ctx, id); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success delete shipping address")
}
