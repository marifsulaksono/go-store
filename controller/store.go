package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/middleware"
	"gostore/service"
	"net/http"
	"time"
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
	result, err := s.Service.GetAllStore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, "Success get all stores")
}

func (s *StoreController) GetStoreById(w http.ResponseWriter, r *http.Request) {
	if id, t := helper.IdVarsMux(w, r); t {
		result, err := s.Service.GetStoreById(id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get store %d", id)
		helper.ResponseWrite(w, result, message)
	}
}

func (s *StoreController) CreateStore(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)
	var store entity.Store
	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	store.Status = "active"
	store.UserId = userId
	store.CreateAt = time.Now()
	if err := s.Service.CreateStore(&store); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Success create new store on user %d", userId)
	helper.ResponseWrite(w, store, message)
}

func (s *StoreController) UpdateStore(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)
	if id, t := helper.IdVarsMux(w, r); t {
		var store entity.Store
		if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err := s.Service.UpdateStore(userId, id, &store)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success update store %d", id)
		helper.ResponseWrite(w, store, message)
	}
}

func (s *StoreController) SoftDeleteStore(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)
	if id, t := helper.IdVarsMux(w, r); t {
		err := s.Service.SoftDeleteStore(userId, id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success delete store %d on user %d", id, userId)
		helper.ResponseWrite(w, id, message)
	}
}

func (s *StoreController) RestoreDeletedStore(w http.ResponseWriter, r *http.Request) {
	if id, t := helper.IdVarsMux(w, r); t {
		err := s.Service.RestoreDeletedStore(id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success restore store %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (s *StoreController) DeleteStore(w http.ResponseWriter, r *http.Request) {
	if id, t := helper.IdVarsMux(w, r); t {
		err := s.Service.DeleteStore(id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success delete store %d", id)
		helper.ResponseWrite(w, id, message)
	}
}
