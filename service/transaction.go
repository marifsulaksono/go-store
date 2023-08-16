package service

import (
	"context"
	"gostore/entity"
	"gostore/helper"
	"gostore/middleware"
	"gostore/repo"
)

type transactionService struct {
	Repo        repo.TransactionRepository
	ProductRepo repo.ProductRepository
	UserRepo    repo.UserRepository
}

type TransactionService interface {
	GetTransactions(ctx context.Context) ([]entity.AllTransactionResponse, error)
	GetTransactionById(ctx context.Context, id int) (entity.Transaction, error)
	CreateTransaction(ctx context.Context, items *entity.Transaction) (entity.Transaction, error)
	SoftDeleteTransaction(ctx context.Context, id, user int) error
	RestoreDeletedTransaction(ctx context.Context, id int) error
	DeleteTransaction(ctx context.Context, id int) error
}

func NewTransactionService(r repo.TransactionRepository, p repo.ProductRepository, u repo.UserRepository) TransactionService {
	return &transactionService{
		Repo:        r,
		ProductRepo: p,
		UserRepo:    u,
	}
}

func (tr *transactionService) GetTransactions(ctx context.Context) ([]entity.AllTransactionResponse, error) {
	return tr.Repo.GetTransactions(ctx)
}

func (tr *transactionService) GetTransactionById(ctx context.Context, id int) (entity.Transaction, error) {
	return tr.Repo.GetTransactionById(ctx, id)
}

func (tr *transactionService) CreateTransaction(ctx context.Context, items *entity.Transaction) (entity.Transaction, error) {
	checkSA, err := tr.UserRepo.GetShippingAddressById(ctx, items.ShippingAddressId)
	if checkSA.UserId != ctx.Value(middleware.GOSTORE_USERID).(int) {
		return entity.Transaction{}, helper.ErrInvalidSA
	} else if err != nil {
		return entity.Transaction{}, err
	}

	return tr.Repo.CreateTransaction(ctx, items)
}

func (tr *transactionService) SoftDeleteTransaction(ctx context.Context, id, user int) error {
	trx, err := tr.Repo.GetTransactionById(ctx, id)
	if err != nil {
		return helper.ErrRecDeleted
	} else if trx.Id != user {
		return helper.ErrAccDeny
	}

	return tr.Repo.SoftDeleteTransaction(ctx, id)
}

func (tr *transactionService) RestoreDeletedTransaction(ctx context.Context, id int) error {
	_, err := tr.Repo.GetTransactionById(ctx, id)
	if err == nil {
		return helper.ErrRecRestored
	}

	return tr.Repo.RestoreDeletedTransaction(ctx, id)
}

func (tr *transactionService) DeleteTransaction(ctx context.Context, id int) error {
	return tr.Repo.DeleteTransaction(ctx, id)
}
