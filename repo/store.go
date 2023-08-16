package repo

import (
	"context"
	"gostore/entity"

	"gorm.io/gorm"
)

type storeRepository struct {
	DB *gorm.DB
}

type StoreRepository interface {
	GetAllStore(ctx context.Context) ([]entity.Store, error)
	GetStoreById(ctx context.Context, id int) (entity.StoreResponseById, error)
	CheckStoreById(ctx context.Context, id int) (entity.Store, error)
	GetDeletedStore(ctx context.Context, id int) (entity.Store, error)
	CreateStore(ctx context.Context, store *entity.Store) error
	UpdateStore(ctx context.Context, id int, model, store *entity.Store) error
	SoftDeleteStore(ctx context.Context, id int) error
	RestoreDeletedStore(ctx context.Context, id int) error
	DeleteStore(ctx context.Context, id int) error
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{
		DB: db,
	}
}

func (s *storeRepository) GetAllStore(ctx context.Context) ([]entity.Store, error) {
	var result []entity.Store
	err := s.DB.Find(&result).Error
	return result, err
}

func (s *storeRepository) GetStoreById(ctx context.Context, id int) (entity.StoreResponseById, error) {
	var result entity.StoreResponseById
	err := s.DB.Where("id = ?", id).Preload("Product.Category").First(&result).Error
	return result, err
}

func (s *storeRepository) CheckStoreById(ctx context.Context, id int) (entity.Store, error) {
	var result entity.Store
	err := s.DB.Where("id = ?", id).First(&result).Error
	return result, err
}

func (s *storeRepository) GetDeletedStore(ctx context.Context, id int) (entity.Store, error) {
	var result entity.Store
	err := s.DB.Unscoped().Where("id = ?", id).First(&result).Error
	return result, err
}

func (s *storeRepository) CreateStore(ctx context.Context, store *entity.Store) error {
	return s.DB.Create(&store).Error
}

func (s *storeRepository) UpdateStore(ctx context.Context, id int, model, store *entity.Store) error {
	return s.DB.Model(model).Where("id = ?", id).Updates(store).Error
}

func (s *storeRepository) SoftDeleteStore(ctx context.Context, id int) error {
	return s.DB.Where("id = ?", id).Delete(&entity.Store{}).Error
}

func (s *storeRepository) RestoreDeletedStore(ctx context.Context, id int) error {
	return s.DB.Unscoped().Model(&entity.Store{}).Where("id = ?", id).Update("delete_at", gorm.DeletedAt{}).Error
}

func (s *storeRepository) DeleteStore(ctx context.Context, id int) error {
	return s.DB.Unscoped().Where("id = ?", id).Delete(&entity.Store{}).Error
}
