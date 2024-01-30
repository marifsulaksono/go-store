package service

import (
	"context"
	"gostore/entity"
	"gostore/repo"
	"gostore/utils/helper"
	transactionError "gostore/utils/helper/domain/errorModel"
)

type transactionService struct {
	Repo   repo.TransactionRepository
	SARepo repo.ShippingAddressRepo
}

type TransactionService interface {
	GetTransactions(ctx context.Context) ([]entity.AllTransactionResponse, error)
	GetTransactionById(ctx context.Context, id string) (entity.Transaction, error)
	CreateTransaction(ctx context.Context, items *entity.Transaction) error
}

func NewTransactionService(r repo.TransactionRepository, sa repo.ShippingAddressRepo) TransactionService {
	return &transactionService{
		Repo:   r,
		SARepo: sa,
	}
}

func (tr *transactionService) GetTransactions(ctx context.Context) ([]entity.AllTransactionResponse, error) {
	return tr.Repo.GetTransactions(ctx)
}

func (tr *transactionService) GetTransactionById(ctx context.Context, id string) (entity.Transaction, error) {
	return tr.Repo.GetTransactionById(ctx, id)
}

func (tr *transactionService) CreateTransaction(ctx context.Context, items *entity.Transaction) error {
	if items.ShippingAddressId == nil {
		detailError := map[string]any{"shipping_address_id": "this field is missing input"}
		return transactionError.ErrTransactionInput.AttachDetail(detailError)
	}

	checkSA, err := tr.SARepo.GetShippingAddressById(ctx, *items.ShippingAddressId)
	if checkSA.UserId != ctx.Value(helper.GOSTORE_USERID).(int) {
		return transactionError.ErrInvalidSA
	} else if err != nil {
		return err
	}

	return tr.Repo.CreateTransaction(ctx, items)
}
