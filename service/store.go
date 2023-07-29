package service

import (
	_ "errors"
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"
	"strings"
)

type StoreService struct {
	Repo repo.StoreRepository
}

func NewStoreService(r repo.StoreRepository) *StoreService {
	return &StoreService{
		Repo: r,
	}
}

func (s *StoreService) GetAllStore() ([]entity.Store, error) {
	result, err := s.Repo.GetAllStore()
	return result, err
}

func (s *StoreService) GetStoreById(id int) (entity.Store, error) {
	result, err := s.Repo.GetStoreById(id)
	return result, err
}

func (s *StoreService) CreateStore(store *entity.Store) error {
	err := s.Repo.CreateStore(store)
	if strings.Contains(err.Error(), "Error 1062") {
		return helper.ErrDuplicateStore
	}
	return err
}

func (s *StoreService) UpdateStore(userId, id int, store *entity.Store) error {
	checkStore, err := s.Repo.GetStoreById(id)
	if checkStore.UserId != userId {
		return helper.ErrInvalidUser
	} else if err != nil {
		return err
	}

	err = s.Repo.UpdateStore(id, &checkStore, store)
	return err
}

func (s *StoreService) SoftDeleteStore(userId, id int) error {
	checkStore, err := s.Repo.GetStoreById(id)
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
