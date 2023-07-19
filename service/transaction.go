package service

import (
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"
)

type TransactionService struct {
	Repo repo.TransactionRepository
}

func NewTransactionService(r repo.TransactionRepository) *TransactionService {
	return &TransactionService{
		Repo: r,
	}
}

func (tr *TransactionService) GetTransactions() ([]entity.AllTransactionResponse, error) {
	result, err := tr.Repo.GetTransactions()
	return result, err
}

func (tr *TransactionService) GetTransactionById(id int) (entity.TransactionResponseId, error) {
	result, err := tr.Repo.GetTransactionById(id)
	return result, err
}

func (tr *TransactionService) CreateTransaction(transaction *entity.Transaction) error {
	err := tr.Repo.CreateTransaction(transaction)
	return err
}

func (tr *TransactionService) UpdateTransaction(id int, transaction *entity.Transaction, userId int) error {
	trxId, err := tr.Repo.GetTransactionById(id)
	if trxId.UserId != userId {
		return helper.ErrAccDeny
	} else if err == nil {
		err = tr.Repo.UpdateTransaction(id, transaction)
	}

	return err
}

func (tr *TransactionService) SoftDeleteTransaction(id int, user int) error {
	trx, err := tr.Repo.GetTransactionById(id)
	if err != nil {
		return helper.ErrRecDeleted
	} else if trx.Id != user {
		return helper.ErrAccDeny
	}

	err = tr.Repo.SoftDeleteTransaction(id)
	return err
}

func (tr *TransactionService) RestoreDeletedTransaction(id int) error {
	_, err := tr.Repo.GetTransactionById(id)
	if err == nil {
		return helper.ErrRecRestored
	}

	err = tr.Repo.RestoreDeletedTransaction(id)
	return err
}

func (tr *TransactionService) DeleteTransaction(id int) error {
	err := tr.Repo.DeleteTransaction(id)
	return err
}
