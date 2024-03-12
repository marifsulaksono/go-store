package service

import (
	"context"
	"gostore/entity"
	"gostore/repo"
	"gostore/utils/helper"
	transactionError "gostore/utils/helper/domain/errorModel"
)

type transactionService struct {
	Repo      repo.TransactionRepository
	SARepo    repo.ShippingAddressRepo
	NotifRepo repo.NotificationRepository
}

type TransactionService interface {
	GetTransactions(ctx context.Context) ([]entity.AllTransactionResponse, error)
	GetTransactionById(ctx context.Context, id string) (entity.Transaction, error)
	CreateTransaction(ctx context.Context, items *entity.Transaction) error
}

func NewTransactionService(r repo.TransactionRepository, sa repo.ShippingAddressRepo, n repo.NotificationRepository) TransactionService {
	return &transactionService{
		Repo:      r,
		SARepo:    sa,
		NotifRepo: n,
	}
}

func (tr *transactionService) GetTransactions(ctx context.Context) ([]entity.AllTransactionResponse, error) {
	return tr.Repo.GetTransactions(ctx)
}

func (tr *transactionService) GetTransactionById(ctx context.Context, id string) (entity.Transaction, error) {
	return tr.Repo.GetTransactionById(ctx, id)
}

func (tr *transactionService) CreateTransaction(ctx context.Context, items *entity.Transaction) error {
	userId := ctx.Value(helper.GOSTORE_USERID).(int)

	if items.ShippingAddressId == nil {
		detailError := map[string]any{"shipping_address_id": "this field is missing input"}
		return transactionError.ErrTransactionInput.AttachDetail(detailError)
	}

	checkSA, err := tr.SARepo.GetShippingAddressById(ctx, *items.ShippingAddressId)
	if checkSA.UserId != userId {
		return transactionError.ErrInvalidSA
	} else if err != nil {
		return err
	}

	transaction, err := tr.Repo.CreateTransaction(ctx, items, userId)
	if err != nil {
		return err
	}

	notif := entity.Notification{
		Title:                  "Transaksi kamu berhasil dibuat!",
		Detail:                 "Hai, transaksi kamu berhasil dibuat, jangan lupa untuk membayarnya ya, klik untuk lebih lanjut. ",
		NotificationCategoryId: 1,
		RedirectUrl:            transaction.PaymentUrl,
		UserId:                 userId,
	}

	err = tr.NotifRepo.InsertNotification(&notif)
	if err != nil {
		return err
	}

	return nil
}
