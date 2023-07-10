package repo

import (
	"gostore/config"
	"gostore/entity"
)

func GetTransactionItem() ([]entity.TransactionItemResponse, error) {
	var result []entity.TransactionItemResponse
	err := config.DB.Preload("Item").Find(&result).Error
	return result, err
}

func GetTransactionItemById(id int64) (entity.TransactionItemResponse, error) {
	var result entity.TransactionItemResponse
	err := config.DB.Where("id = ?", id).Preload("Item").First(&result).Error
	return result, err
}

func CreateTransactionItem(transactionItem entity.TransactionItem) error {
	err := config.DB.Create(&transactionItem).Error
	return err
}

func UpdateTransactionItem(id int64, transactionItem entity.TransactionItem) error {
	err := config.DB.Model(&entity.TransactionItem{}).Where("id = ?", id).Updates(&transactionItem).Error
	return err
}

func DeleteTransactionItem(id int64) error {
	err := config.DB.Where("id = ?", id).Delete(&entity.TransactionItem{}).Error
	return err
}
