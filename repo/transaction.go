package repo

import (
	"context"
	"errors"
	"fmt"
	"gostore/entity"
	transactionError "gostore/helper/domain/errorModel"
	"gostore/middleware"
	"time"

	"gorm.io/gorm"
)

type transactionRepository struct {
	DB *gorm.DB
}

// TransactionRepository: represent the transactionRepository contract
type TransactionRepository interface {
	GetTransactions(ctx context.Context) ([]entity.AllTransactionResponse, error)
	GetTransactionById(ctx context.Context, id int) (entity.Transaction, error)
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) error
}

// return new transaction repository with property value
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		DB: db,
	}
}

func (tr *transactionRepository) GetTransactions(ctx context.Context) ([]entity.AllTransactionResponse, error) {
	var (
		result []entity.AllTransactionResponse
		userId = ctx.Value(middleware.GOSTORE_USERID)
	)

	err := tr.DB.Where("user_id = ?", userId).Preload("Items.Product.Store").Find(&result).Error
	return result, err
}

func (tr *transactionRepository) GetTransactionById(ctx context.Context, id int) (entity.Transaction, error) {
	var result entity.Transaction
	err := tr.DB.Where("id = ?", id).Preload("Items.Product.Store").Preload("ShippingAddress").First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Transaction{}, transactionError.ErrTransactionNotFound
		}
		return entity.Transaction{}, err
	}

	return result, nil
}

func (tr *transactionRepository) CreateTransaction(ctx context.Context, items *entity.Transaction) error {
	var (
		transaction entity.Transaction
		total       int
		userId      = ctx.Value(middleware.GOSTORE_USERID).(int)
		detailError = make(map[string]any)
	)

	// begin the database transaction for ACID (atomicity, Consistency, Isolation, Durability)
	tx := tr.DB.Begin()

	// validation and update stock, sold, status selected product
	for i, trItem := range items.Items {
		if trItem.ProductId == nil {
			detailError["product_id"] = "this field is missing input"
		}

		if trItem.Qty == nil {
			detailError["qty"] = "this field is missing input"
		}

		if len(detailError) > 0 {
			tx.Rollback()
			return transactionError.ErrTransactionInput.AttachDetail(detailError)
		}

		// check product available
		var product entity.Product
		err := tx.Where("id = ?", *trItem.ProductId).Preload("Store").First(&product).Error
		if err != nil {
			tx.Rollback()
			detailError["item"] = fmt.Sprintf("Product %d not found", *trItem.ProductId)
			return transactionError.ErrProductNotFound.AttachDetail(detailError)
		} else if product.Store.UserId == userId {
			// can't add user's product to transaction
			tx.Rollback()
			return transactionError.ErrCantAddToTrx
		} else if *product.Stock < *trItem.Qty {
			tx.Rollback()
			return transactionError.ErrStockProductNotEnough
		}

		*product.Stock -= *trItem.Qty
		product.Sold += *trItem.Qty
		trItem.Subtotal = *product.Price * *trItem.Qty
		total += trItem.Subtotal
		items.Items[i].Price = *product.Price
		items.Items[i].Subtotal = trItem.Subtotal

		// update stock and sold product
		err = tx.Model(&entity.Product{}).Select("stock", "sold").Where("id = ?",
			*trItem.ProductId).Updates(product).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		if *product.Stock == 0 {
			// update status product when stock empty after transaction
			err := tx.Model(&entity.Product{}).Where("id = ?", *trItem.ProductId).Update("status", "soldout").Error
			if err != nil {
				tx.Rollback()
				return err
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

	err := tx.Create(&transaction).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
