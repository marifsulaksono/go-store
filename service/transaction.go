package service

import (
	"errors"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"
	"time"
)

type TransactionService struct {
	Repo        repo.TransactionRepository
	ProductRepo repo.ProductRepository
	UserRepo    repo.UserRepository
}

func NewTransactionService(r repo.TransactionRepository, p repo.ProductRepository, u repo.UserRepository) *TransactionService {
	return &TransactionService{
		Repo:        r,
		ProductRepo: p,
		UserRepo:    u,
	}
}

func (tr *TransactionService) GetTransactions() ([]entity.AllTransactionResponse, error) {
	result, err := tr.Repo.GetTransactions()
	return result, err
}

func (tr *TransactionService) GetTransactionById(id int) (entity.Transaction, error) {
	result, err := tr.Repo.GetTransactionById(id)
	return result, err
}

func (tr *TransactionService) CreateTransaction(userId int, items *entity.Transaction) (entity.Transaction, error) {
	var transaction entity.Transaction
	var total int
	// check valid shipping address with the user login
	checkSA, err := tr.UserRepo.GetShippingAddressById(items.ShippingAddressId)
	if checkSA.UserId != userId {
		return entity.Transaction{}, helper.ErrInvalidUser
	} else if err != nil {
		return entity.Transaction{}, err
	}
	// validation and update stock, sold, status selected product
	for i, trItem := range items.Items {
		product, err := tr.ProductRepo.GetProductById(trItem.ProductId)
		if err != nil {
			productNotFound := fmt.Sprintf("Product %d not found!", trItem.ProductId)
			return entity.Transaction{}, errors.New(productNotFound)
		} else if product.Stock < trItem.Qty {
			return entity.Transaction{}, helper.ErrStockNotEnough
		}
		product.Stock -= trItem.Qty
		product.Sold += trItem.Qty
		trItem.Subtotal = product.Price * trItem.Qty
		total += trItem.Subtotal
		items.Items[i].Price = product.Price
		items.Items[i].Subtotal = trItem.Subtotal
		fmt.Println(product)
		if err := tr.ProductRepo.UpdateStockandSold(trItem.ProductId, &product); err != nil {
			return entity.Transaction{}, err
		}

		if product.Stock == 0 {
			err := tr.ProductRepo.ChangeStatusProduct(trItem.ProductId, "soldout")
			if err != nil {
				return entity.Transaction{}, err
			}
		}
	}

	// default transaction property
	transaction.Date = time.Now()
	transaction.Total = total
	transaction.Status = "waiting"
	transaction.ShippingAddressId = items.ShippingAddressId
	transaction.UserId = userId
	transaction.Items = items.Items

	err = tr.Repo.CreateTransaction(&transaction)
	return transaction, err
}

// func (tr *TransactionService) UpdateTransaction(id int, transaction *entity.Transaction, userId int) error {
// 	trxId, err := tr.Repo.GetTransactionById(id)
// 	if trxId.UserId != userId {
// 		return helper.ErrAccDeny
// 	} else if err != nil {
// 		return err
// 	}

// 	transaction.UpdateAt = time.Now()
// 	err = tr.Repo.UpdateTransaction(id, transaction)
// 	return err
// }

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
