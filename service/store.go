package service

import (
	"context"
	_ "errors"
	"gostore/entity"
	"gostore/repo"
	"gostore/utils/helper"
	storeError "gostore/utils/helper/domain/errorModel"
	"strings"
	"time"
)

type storeService struct {
	Repo        repo.StoreRepository
	ProductRepo repo.ProductRepository
	UserRepo    repo.UserRepository
}

type StoreService interface {
	GetAllStore(ctx context.Context) ([]entity.Store, error)
	GetStoreById(ctx context.Context, id int) (entity.StoreResponseById, error)
	CreateStore(ctx context.Context, store *entity.Store) error
	UpdateStore(ctx context.Context, id int, store *entity.Store) error
	SoftDeleteStore(ctx context.Context, id int) error
	RestoreDeletedStore(ctx context.Context, id int) error
	DeleteStore(ctx context.Context, id int) error
}

func NewStoreService(r repo.StoreRepository, p repo.ProductRepository, u repo.UserRepository) StoreService {
	return &storeService{
		Repo:        r,
		ProductRepo: p,
		UserRepo:    u,
	}
}

func (s *storeService) GetAllStore(ctx context.Context) ([]entity.Store, error) {
	return s.Repo.GetAllStore(ctx)
}

func (s *storeService) GetStoreById(ctx context.Context, id int) (entity.StoreResponseById, error) {
	result, err := s.Repo.GetStoreById(ctx, id)
	result.TotalProduct = len(result.Product)
	return result, err
}

func (s *storeService) CreateStore(ctx context.Context, store *entity.Store) error {
	var (
		detailError = make(map[string]any)
		userId      = ctx.Value(helper.GOSTORE_USERID).(int)
	)

	if store.NameStore == "" {
		detailError["name_store"] = "this field is missing input"
	}

	if store.Address == "" {
		detailError["address"] = "this field is missing input"
	}

	if store.Email == "" {
		detailError["email"] = "this field is missing input"
	}

	if len(detailError) > 0 {
		return storeError.ErrStoreInput.AttachDetail(detailError)
	}

	store.Status = "active"
	store.UserId = userId
	store.CreateAt = time.Now()
	err := s.Repo.CreateStore(ctx, store)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			return storeError.ErrDuplicateStore
		}
		return err
	}

	err = s.UserRepo.ChangeUserRole(ctx, userId, "seller")
	if err != nil {
		return err
	}

	return nil
}

func (s *storeService) UpdateStore(ctx context.Context, id int, store *entity.Store) error {
	var (
		detailError = make(map[string]any)
		userId      = ctx.Value(helper.GOSTORE_USERID).(int)
	)

	if store.NameStore == "" {
		detailError["name_store"] = "this field is missing input"
	}

	if store.Address == "" {
		detailError["address"] = "this field is missing input"
	}

	if store.Email == "" {
		detailError["email"] = "this field is missing input"
	}

	if store.Status == "" {
		detailError["status"] = "this field is missing input"
	}

	if len(detailError) > 0 {
		return storeError.ErrStoreInput.AttachDetail(detailError)
	}

	checkStore, err := s.Repo.CheckStoreById(ctx, id)
	if checkStore.UserId != userId {
		return storeError.ErrInvalidUserStore
	} else if err != nil {
		return err
	}

	return s.Repo.UpdateStore(ctx, id, &checkStore, store)
}

func (s *storeService) SoftDeleteStore(ctx context.Context, id int) error {
	userId := ctx.Value(helper.GOSTORE_USERID).(int)
	checkStore, err := s.Repo.CheckStoreById(ctx, id)
	if checkStore.UserId != userId {
		return storeError.ErrInvalidUserStore
	} else if err != nil {
		return err
	}

	return s.Repo.SoftDeleteStore(ctx, id)
}

func (s *storeService) RestoreDeletedStore(ctx context.Context, id int) error {
	checkStore, err := s.Repo.GetDeletedStore(ctx, id)
	if err != nil {
		return err
	} else if !checkStore.DeleteAt.Valid {
		return storeError.ErrStoreRestored
	}

	return s.Repo.RestoreDeletedStore(ctx, id)
}

func (s *storeService) DeleteStore(ctx context.Context, id int) error {
	return s.Repo.DeleteStore(ctx, id)
}
