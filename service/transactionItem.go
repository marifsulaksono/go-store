package service

import (
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"
)

type TransactionItemService struct {
	Repo repo.TransactionItemRepository
}

func NewTransactionItemService(r repo.TransactionItemRepository) *TransactionItemService {
	return &TransactionItemService{
		Repo: r,
	}
}

func (ti *TransactionItemService) GetTransactionItems() ([]entity.TransactionItemResponse, error) {
	result, err := ti.Repo.GetTransactionItems()
	return result, err
}

func (ti *TransactionItemService) GetTransactionItemById(id int) (entity.TransactionItemResponse, error) {
	result, err := ti.Repo.GetTransactionItemById(id)
	return result, err
}

func (ti *TransactionItemService) CreateTransactionItem(transactionItem *entity.TransactionItem) error {
	err := ti.Repo.CreateTransactionItem(transactionItem)
	return err
}

func (ti *TransactionItemService) UpdateTransactionItem(id int, transactionItem *entity.TransactionItem) error {
	err := ti.Repo.UpdateTransactionItem(id, transactionItem)
	return err
}

func (ti *TransactionItemService) SoftDeleteTransactionItem(id int) error {
	_, err := ti.Repo.GetTransactionItemById(id)
	if err != nil {
		return helper.ErrRecDeleted
	}

	err = ti.Repo.SoftDeleteTransactionItem(id)
	return err
}

func (ti *TransactionItemService) RestoreDeletedTransactionItem(id int) error {
	_, err := ti.Repo.GetTransactionItemById(id)
	if err == nil {
		return helper.ErrRecRestored
	}

	err = ti.Repo.RestoreDeletedTransactionItem(id)
	return err
}

func (ti *TransactionItemService) DeleteTransactionItem(id int) error {
	err := ti.Repo.DeleteTransactionItem(id)
	return err
}
