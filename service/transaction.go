package service

import (
	"gostore/entity"
	"gostore/repo"
)

func GetTransaction() ([]entity.AllTransactionResponse, error) {
	result, err := repo.GetTransaction()
	return result, err
}

func GetTransactionById(id int64) (entity.TransactionResponseId, error) {
	result, err := repo.GetTransactionById(id)
	return result, err
}

func CreateTransaction(transaction entity.Transaction) error {
	err := repo.CreateTransaction(transaction)
	return err
}

func UpdateTransaction(id int64, transaction entity.Transaction) error {
	err := repo.UpdateTransaction(id, transaction)
	return err
}

func DeleteTransaction(id int64) error {
	err := repo.DeleteTransaction(id)
	return err
}
