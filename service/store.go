package service

import (
	_ "errors"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"
	"strings"
	"time"
)

type StoreService struct {
	Repo        repo.StoreRepository
	ProductRepo repo.ProductRepository
}

func NewStoreService(r repo.StoreRepository, p repo.ProductRepository) *StoreService {
	return &StoreService{
		Repo:        r,
		ProductRepo: p,
	}
}

func (s *StoreService) GetAllStore() ([]entity.Store, error) {
	result, err := s.Repo.GetAllStore()
	return result, err
}

func (s *StoreService) GetStoreById(id int) (entity.StoreResponseById, error) {
	product, err := s.ProductRepo.GetProductOnStore(id)
	if err != nil {
		return entity.StoreResponseById{}, err
	}
	result, err := s.Repo.GetStoreById(id)
	result.TotalProduct = len(product)
	result.Product = append(result.Product, product...)
	return result, err
}

func (s *StoreService) CreateStore(userId int, store *entity.Store) (entity.Store, error) {
	store.Status = "active"
	store.UserId = userId
	store.CreateAt = time.Now()
	err := s.Repo.CreateStore(store)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			fmt.Println("error service disini", err)
			return entity.Store{}, helper.ErrDuplicateStore
		}
		return entity.Store{}, err
	}
	return *store, nil
}

func (s *StoreService) UpdateStore(userId, id int, store *entity.Store) error {
	checkStore, err := s.Repo.CheckStoreById(id)
	if checkStore.UserId != userId {
		return helper.ErrInvalidUser
	} else if err != nil {
		return err
	}

	err = s.Repo.UpdateStore(id, &checkStore, store)
	return err
}

func (s *StoreService) SoftDeleteStore(userId, id int) error {
	checkStore, err := s.Repo.CheckStoreById(id)
	if checkStore.UserId != userId {
		return helper.ErrInvalidUser
	} else if err != nil {
		return err
	}

	err = s.Repo.SoftDeleteStore(id)
	return err
}

func (s *StoreService) RestoreDeletedStore(id int) error {
	checkStore, err := s.Repo.GetDeletedStore(id)
	if err != nil {
		return err
	} else if !checkStore.DeleteAt.Valid {
		return helper.ErrRecRestored
	}

	err = s.Repo.RestoreDeletedStore(id)
	return err
}

func (s *StoreService) DeleteStore(id int) error {
	err := s.Repo.DeleteStore(id)
	return err
}
