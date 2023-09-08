package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/helper/response"
	"gostore/service"
	"net/http"
)

type StoreController struct {
	Service service.StoreService
}

func NewStoreController(s service.StoreService) *StoreController {
	return &StoreController{
		Service: s,
	}
}

func (s *StoreController) GetAllStore(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		message string
	)

	result, err := s.Service.GetAllStore(ctx)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if result == nil || len(result) < 1 {
		message = "No results found"
	}

	response.BuildSuccesResponse(w, result, nil, message)
}

func (s *StoreController) GetStoreById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	result, err := s.Service.GetStoreById(ctx, id)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, result, nil, "")
}

func (s *StoreController) CreateStore(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = r.Context()
		store entity.Store
	)

	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	if err := s.Service.CreateStore(ctx, &store); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success create new store")
}

func (s *StoreController) UpdateStore(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = r.Context()
		store entity.Store
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	if err := s.Service.UpdateStore(ctx, id, &store); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success update store")
}

func (s *StoreController) SoftDeleteStore(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := s.Service.SoftDeleteStore(ctx, id); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success delete store")
}

func (s *StoreController) RestoreDeletedStore(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := s.Service.RestoreDeletedStore(ctx, id); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	message := fmt.Sprintf("Success restore store %d", id)
	response.BuildSuccesResponse(w, nil, nil, message)
}

func (s *StoreController) DeleteStore(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := s.Service.DeleteStore(ctx, id); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	message := fmt.Sprintf("Success permanently delete store %d", id)
	response.BuildSuccesResponse(w, nil, nil, message)
}
