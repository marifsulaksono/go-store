package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id                string            `gorm:"primaryKey" json:"id"`
	Date              time.Time         `gorm:"not null" json:"date"`
	Total             int               `gorm:"not null" json:"total"`
	Status            string            `gorm:"not null" json:"status"`
	ShippingAddressId *int              `gorm:"not null" json:"shipping_address_id"`
	ShippingAddress   ShippingAddress   `gorm:"-:migration" json:"shipping_address"`
	PaymentUrl        string            `json:"payment_url"`
	UserId            int               `gorm:"not null" json:"-"`
	Items             []TransactionItem `json:"items" gorm:"-:migration;foreignKey:TransactionId;references:Id"`
	DeleteAt          gorm.DeletedAt    `json:"-"`
}

type TransactionItem struct {
	Id            int                        `gorm:"primaryKey,autoIncrement" json:"id"`
	TransactionId string                     `gorm:"not null" json:"-"`
	ProductId     *int                       `gorm:"not null" json:"product_id"`
	Product       ProductTransactionResponse `gorm:"-:migration" json:"items"`
	Qty           *int                       `gorm:"not null" json:"qty"`
	Price         int                        `gorm:"not null" json:"price"`
	Subtotal      int                        `gorm:"not null" json:"subtotal"`
}

type AllTransactionResponse struct {
	Id       string                       `json:"id"`
	Date     time.Time                    `json:"date"`
	Total    int                          `json:"total"`
	Status   string                       `json:"status"`
	Items    []AllTransactionItemResponse `json:"transaction_items" gorm:"foreignKey:TransactionId;references:Id"`
	DeleteAt gorm.DeletedAt               `json:"-"`
}

type AllTransactionItemResponse struct {
	Id            int                        `json:"id"`
	TransactionId string                     `json:"-"`
	ProductId     int                        `json:"product_id"`
	Product       ProductTransactionResponse `gorm:"foreignKey:ProductId" json:"product"`
	Subtotal      int                        `json:"subtotal"`
}

func (AllTransactionResponse) TableName() string {
	return "transactions"
}

func (AllTransactionItemResponse) TableName() string {
	return "transaction_items"
}

// {
//     "shipping_address_id": 1,
//     "items": [
//         {
//             "product_id": 1,
//             "qty": 4
//         },
//         {
//             "product_id": 19,
//             "qty": 4
//         }
//     ]
// }
