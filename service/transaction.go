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
	CreateTransaction(ctx context.Context, items *entity.Transaction) error
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

func (tr *transactionService) CreateTransaction(ctx context.Context, items *entity.Transaction) error {
	checkSA, err := tr.UserRepo.GetShippingAddressById(ctx, items.ShippingAddressId)
	if checkSA.UserId != ctx.Value(middleware.GOSTORE_USERID).(int) {
		return helper.ErrInvalidSA
	} else if err != nil {
		return err
	}

	return tr.Repo.CreateTransaction(ctx, items)
}
