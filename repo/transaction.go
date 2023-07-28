package repo

import (
	"gostore/entity"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (tr *TransactionRepository) GetTransactions() ([]entity.AllTransactionResponse, error) {
	var result []entity.AllTransactionResponse
	err := tr.DB.Preload("Items").Find(&result).Error
	return result, err
}

func (tr *TransactionRepository) GetTransactionById(id int) (entity.Transaction, error) {
	var result entity.Transaction
	err := tr.DB.Where("id = ?", id).Preload("Items.Product").Preload("Items.Product.Category").Preload("ShippingAddress").First(&result).Error
	return result, err
}

func (tr *TransactionRepository) CreateTransaction(transaction *entity.Transaction) error {
	err := tr.DB.Create(transaction).Error
	return err
}

// func (tr *TransactionRepository) UpdateTransaction(id int, transaction *entity.Transaction) error {
// 	err := tr.DB.Model(entity.Transaction{}).Where("id = ?", id).Updates(map[string]interface{}{
// 		"total":     transaction.Total,
// 		"status":    transaction.Status,
// 		"update_at": transaction.UpdateAt,
// 	}).Error
// 	return err
// }

func (tr *TransactionRepository) SoftDeleteTransaction(id int) error {
	err := tr.DB.Where("id = ?", id).Delete(&entity.Transaction{}).Error
	return err
}

func (tr *TransactionRepository) RestoreDeletedTransaction(id int) error {
	err := tr.DB.Unscoped().Model(&entity.Transaction{}).Where("id = ?", id).Update("delete_at", gorm.DeletedAt{}).Error
	return err
}

func (tr *TransactionRepository) DeleteTransaction(id int) error {
	err := tr.DB.Unscoped().Where("id = ?", id).Delete(&entity.Transaction{}).Error
	return err
}
