package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id                int               `json:"id"`
	Date              time.Time         `json:"date"`
	Total             int               `json:"total"`
	Status            string            `json:"status"`
	ShippingAddressId int               `json:"shipping_address_id"`
	ShippingAddress   ShippingAddress   `json:"shipping_address"`
	UserId            int               `json:"user_id"`
	Items             []TransactionItem `json:"items" gorm:"foreignKey:TransactionId;references:Id"`
	DeleteAt          gorm.DeletedAt    `json:"-"`
}

type TransactionItem struct {
	Id            int                        `json:"id"`
	TransactionId int                        `json:"-"`
	ProductId     int                        `json:"product_id"`
	Product       ProductTransactionResponse `gorm:"foreignKey:ProductId" json:"items"`
	Qty           int                        `json:"qty"`
	Price         int                        `json:"price"`
	Subtotal      int                        `json:"subtotal"`
}

type AllTransactionResponse struct {
	Id       int                          `json:"id"`
	Date     time.Time                    `json:"date"`
	Total    int                          `json:"total"`
	Status   string                       `json:"status"`
	Items    []AllTransactionItemResponse `json:"transaction_items" gorm:"foreignKey:TransactionId;references:Id"`
	DeleteAt gorm.DeletedAt               `json:"-"`
}

type AllTransactionItemResponse struct {
	Id            int                        `json:"id"`
	TransactionId int                        `json:"-"`
	ProductId     int                        `json:"product_id"`
	Product       ProductTransactionResponse `gorm:"foreignKey:ProductId" json:"items"`
	Subtotal      int                        `json:"subtotal"`
}

func (AllTransactionResponse) TableName() string {
	return "transactions"
}

func (AllTransactionItemResponse) TableName() string {
	return "transaction_items"
}

// {
//     "transaction_items": [
//         {
//             "item_id": 1,
//             "qty": 4
//         },
//         {
//             "item_id": 19,
//             "qty": 4
//         }
//     ]
// }
