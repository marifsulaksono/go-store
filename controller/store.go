package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
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
	ctx := r.Context()
	result, err := s.Service.GetAllStore(ctx)
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

func (s *StoreController) GetStoreById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, t := helper.IdVarsMux(w, r); t {
		result, err := s.Service.GetStoreById(ctx, id)
		if err != nil {
			helper.BuildError(w, err)
			return
		}

		helper.BuildResponseSuccess(w, result, nil, "")
	}
}

func (s *StoreController) CreateStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var store entity.Store
	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		helper.BuildError(w, err)
		return
	}
	defer r.Body.Close()

	err := s.Service.CreateStore(ctx, &store)
	if err != nil {
		helper.BuildError(w, err)
		return
	}

	helper.BuildResponseSuccess(w, nil, nil, "Success create new store")
}

func (s *StoreController) UpdateStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, t := helper.IdVarsMux(w, r); t {
		var store entity.Store
		if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
			helper.BuildError(w, err)
			return
		}

		err := s.Service.UpdateStore(ctx, id, &store)
		if err != nil {
			helper.BuildError(w, err)
			return
		}

		helper.BuildResponseSuccess(w, nil, nil, "Success update store")
	}
}

func (s *StoreController) SoftDeleteStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, t := helper.IdVarsMux(w, r); t {
		err := s.Service.SoftDeleteStore(ctx, id)
		if err != nil {
			helper.BuildError(w, err)
			return
		}

		helper.BuildResponseSuccess(w, nil, nil, "Success delete store")
	}
}

func (s *StoreController) RestoreDeletedStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, t := helper.IdVarsMux(w, r); t {
		err := s.Service.RestoreDeletedStore(ctx, id)
		if err != nil {
			helper.BuildError(w, err)
			return
		}

		message := fmt.Sprintf("Success restore store %d", id)
		helper.BuildResponseSuccess(w, nil, nil, message)
	}
}

func (s *StoreController) DeleteStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, t := helper.IdVarsMux(w, r); t {
		err := s.Service.DeleteStore(ctx, id)
		if err != nil {
			helper.BuildError(w, err)
			return
		}

		message := fmt.Sprintf("Success permanently delete store %d", id)
		helper.BuildResponseSuccess(w, nil, nil, message)
	}
}
