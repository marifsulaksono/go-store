package repo

import (
	"gostore/entity"

	"gorm.io/gorm"
)

type StoreRepository struct {
	DB *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepository {
	return &StoreRepository{
		DB: db,
	}
}

func (s *StoreRepository) GetAllStore() ([]entity.Store, error) {
	var result []entity.Store
	err := s.DB.Find(&result).Error
	return result, err
}

func (s *StoreRepository) GetStoreById(id int) (entity.Store, error) {
	var result entity.Store
	err := s.DB.Where("id = ?", id).First(&result).Error
	return result, err
}

func (s *StoreRepository) GetDeletedStore(id int) (entity.Store, error) {
	var result entity.Store
	err := s.DB.Unscoped().Where("id = ?", id).First(&result).Error
	return result, err
}

func (s *StoreRepository) CreateStore(store *entity.Store) error {
	err := s.DB.Create(store).Error
	return err
}

func (s *StoreRepository) UpdateStore(id int, model, store *entity.Store) error {
	err := s.DB.Model(model).Where("id = ?", id).Updates(store).Error
	return err
}

func (s *StoreRepository) SoftDeleteStore(id int) error {
	err := s.DB.Where("id = ?", id).Delete(&entity.Store{}).Error
	return err
}

func (s *StoreRepository) RestoreDeletedStore(id int) error {
	err := s.DB.Unscoped().Model(&entity.Store{}).Where("id = ?", id).Update("delete_at", gorm.DeletedAt{}).Error
	return err
}

func (s *StoreRepository) DeleteStore(id int) error {
	err := s.DB.Unscoped().Where("id = ?", id).Delete(&entity.Store{}).Error
	return err
}
