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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, "Success get all stores")
}

func (s *StoreController) GetStoreById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, t := helper.IdVarsMux(w, r); t {
		result, err := s.Service.GetStoreById(ctx, id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get store %d", id)
		helper.ResponseWrite(w, result, message)
	}
}

func (s *StoreController) CreateStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var store entity.Store
	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	result, err := s.Service.CreateStore(ctx, &store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, "Success create new store")
}

func (s *StoreController) UpdateStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, t := helper.IdVarsMux(w, r); t {
		var store entity.Store
		if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err := s.Service.UpdateStore(ctx, id, &store)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success update store %d", id)
		helper.ResponseWrite(w, store, message)
	}
}

func (s *StoreController) SoftDeleteStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, t := helper.IdVarsMux(w, r); t {
		err := s.Service.SoftDeleteStore(ctx, id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success delete store %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (s *StoreController) RestoreDeletedStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, t := helper.IdVarsMux(w, r); t {
		err := s.Service.RestoreDeletedStore(ctx, id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success restore store %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (s *StoreController) DeleteStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, t := helper.IdVarsMux(w, r); t {
		err := s.Service.DeleteStore(ctx, id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success delete store %d", id)
		helper.ResponseWrite(w, id, message)
	}
}
