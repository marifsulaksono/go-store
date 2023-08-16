package repo

import (
	"context"
	"errors"
	"fmt"
	"gostore/entity"
	"gostore/helper"
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
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (entity.Transaction, error)
	SoftDeleteTransaction(ctx context.Context, id int) error
	RestoreDeletedTransaction(ctx context.Context, id int) error
	DeleteTransaction(ctx context.Context, id int) error
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
	return result, err
}

func (tr *transactionRepository) CreateTransaction(ctx context.Context, items *entity.Transaction) (entity.Transaction, error) {
	var (
		transaction entity.Transaction
		total       int
		userId      = ctx.Value(middleware.GOSTORE_USERID).(int)
	)

	// begin the database transaction for ACID (atomicity, Consistency, Isolation, Durability)
	tx := tr.DB.Begin()

	// validation and update stock, sold, status selected product
	for i, trItem := range items.Items {
		// check product available
		var product entity.Product
		err := tx.Where("id = ?", trItem.ProductId).Preload("Store").First(&product).Error
		fmt.Println("product Store UserId: ", product.Store.UserId)
		if err != nil {
			tx.Rollback()
			productNotFound := fmt.Sprintf("Product %d not found!", trItem.ProductId)
			return entity.Transaction{}, errors.New(productNotFound)
		} else if product.Store.UserId == userId {
			tx.Rollback()
			// can't add user's product to transaction
			return entity.Transaction{}, helper.ErrAddProductTo
		} else if product.Stock < trItem.Qty {
			tx.Rollback()
			return entity.Transaction{}, helper.ErrStockNotEnough
		}
		fmt.Println("stock awal :", product.Stock)
		product.Stock -= trItem.Qty
		product.Sold += trItem.Qty
		trItem.Subtotal = product.Price * trItem.Qty
		total += trItem.Subtotal
		items.Items[i].Price = product.Price
		items.Items[i].Subtotal = trItem.Subtotal
		fmt.Println(product)

		// update stock and sold product
		err = tx.Model(&entity.Product{}).Select("stock", "sold").Where("id = ?",
			trItem.ProductId).Updates(product).Error
		if err != nil {
			tx.Rollback()
			return entity.Transaction{}, err
		}

		fmt.Println("stock akhir :", product.Stock)
		if product.Stock == 0 {
			// update status product when stock empty after transaction
			err := tx.Model(&entity.Product{}).Where("id = ?", trItem.ProductId).Update("status", "soldout").Error
			if err != nil {
				tx.Rollback()
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

	fmt.Println(transaction.Items)
	err := tx.Create(&transaction).Error
	if err != nil {
		tx.Rollback()
		return entity.Transaction{}, err
	}

	return transaction, tx.Commit().Error
}

func (tr *transactionRepository) SoftDeleteTransaction(ctx context.Context, id int) error {
	return tr.DB.Where("id = ?", id).Delete(&entity.Transaction{}).Error
}

func (tr *transactionRepository) RestoreDeletedTransaction(ctx context.Context, id int) error {
	return tr.DB.Unscoped().Model(&entity.Transaction{}).Where("id = ?", id).Update("delete_at", gorm.DeletedAt{}).Error
}

func (tr *transactionRepository) DeleteTransaction(ctx context.Context, id int) error {
	return tr.DB.Unscoped().Where("id = ?", id).Delete(&entity.Transaction{}).Error
}
