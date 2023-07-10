package repo

import (
	"gostore/config"
	"gostore/entity"
)

func GetTransaction() ([]entity.AllTransactionResponse, error) {
	var result []entity.AllTransactionResponse
	err := config.DB.Preload("TransactionItem").Find(&result).Error
	return result, err
}

func GetTransactionById(id int64) (entity.TransactionResponseId, error) {
	var result entity.TransactionResponseId
	err := config.DB.Where("id = ?", id).Preload("TransactionItem").Preload("TransactionItem.Item").First(&result).Error
	return result, err
}

func CreateTransaction(transaction entity.Transaction) error {
	err := config.DB.Create(&transaction).Error
	return err
}

func UpdateTransaction(id int64, transaction entity.Transaction) error {
	err := config.DB.Model(&entity.Transaction{}).Where("id = ?", id).Updates(transaction).Error
	return err
}

func DeleteTransaction(id int64) error {
	err := config.DB.Where("id = ?", id).Delete(&entity.Transaction{}).Error
	return err
}
