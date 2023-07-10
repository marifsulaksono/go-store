package service

import (
	"gostore/entity"
	"gostore/repo"
)

func GetTransactionItem() ([]entity.TransactionItemResponse, error) {
	result, err := repo.GetTransactionItem()
	return result, err
}

func GetTransactionItemById(id int64) (entity.TransactionItemResponse, error) {
	result, err := repo.GetTransactionItemById(id)
	return result, err
}

func CreateTransactionItem(transactionItem entity.TransactionItem) error {
	err := repo.CreateTransactionItem(transactionItem)
	return err
}

func UpdateTransactionItem(id int64, transactionItem entity.TransactionItem) error {
	err := repo.UpdateTransactionItem(id, transactionItem)
	return err
}

func DeleteTransactionItem(id int64) error {
	err := repo.DeleteTransactionItem(id)
	return err
}
