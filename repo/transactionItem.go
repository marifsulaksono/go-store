package repo

import (
	"gostore/entity"

	"gorm.io/gorm"
)

type TransactionItemRepository struct {
	DB *gorm.DB
}

func NewTransactionItemRepository(db *gorm.DB) *TransactionItemRepository {
	return &TransactionItemRepository{
		DB: db,
	}
}

func (ti *TransactionItemRepository) GetTransactionItems() ([]entity.TransactionItemResponse, error) {
	var result []entity.TransactionItemResponse
	err := ti.DB.Preload("Item").Find(&result).Error
	return result, err
}

func (ti *TransactionItemRepository) GetTransactionItemById(id int) (entity.TransactionItemResponse, error) {
	var result entity.TransactionItemResponse
	err := ti.DB.Where("id = ?", id).Preload("Item").First(&result).Error
	return result, err
}

func (ti *TransactionItemRepository) CreateTransactionItem(transactionItem *entity.TransactionItem) error {
	err := ti.DB.Create(&transactionItem).Error
	return err
}

func (ti *TransactionItemRepository) UpdateTransactionItem(id int, transactionItem *entity.TransactionItem) error {
	err := ti.DB.Model(&entity.TransactionItem{}).Where("id = ?", id).Updates(&transactionItem).Error
	return err
}

func (ti *TransactionItemRepository) SoftDeleteTransactionItem(id int) error {
	err := ti.DB.Where("id = ?", id).Delete(&entity.TransactionItem{}).Error
	return err
}

func (ti *TransactionItemRepository) RestoreDeletedTransactionItem(id int) error {
	err := ti.DB.Unscoped().Model(&entity.TransactionItem{}).Where("id = ?", id).Update("delete_at", gorm.DeletedAt{}).Error
	return err
}

func (ti *TransactionItemRepository) DeleteTransactionItem(id int) error {
	err := ti.DB.Unscoped().Where("id = ?", id).Delete(&entity.TransactionItem{}).Error
	return err
}
