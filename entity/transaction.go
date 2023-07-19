package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id       int            `json:"id"`
	Date     time.Time      `json:"date"`
	Total    int            `json:"total"`
	Status   string         `json:"status"`
	UserId   int            `json:"user_id"`
	UpdateAt time.Time      `json:"update_at"`
	DeleteAt gorm.DeletedAt `json:"delete_at"`
}

type AllTransactionResponse struct {
	Id              int                          `json:"id"`
	Date            time.Time                    `json:"date"`
	Total           int                          `json:"total"`
	Status          string                       `json:"status"`
	UserId          int                          `json:"user_id"`
	TransactionItem []AllTransactionItemResponse `gorm:"ForeignKey:TransactionId" json:"transaction_item"`
	UpdateAt        time.Time                    `json:"update_at"`
	DeleteAt        gorm.DeletedAt               `json:"-"`
}

type TransactionResponseId struct {
	Id              int                       `json:"id"`
	Date            time.Time                 `json:"date"`
	Total           int                       `json:"total"`
	Status          string                    `json:"status"`
	UserId          int                       `json:"user_id"`
	UpdateAt        time.Time                 `json:"update_at"`
	TransactionItem []TransactionItemResponse `gorm:"ForeignKey:TransactionId" json:"transaction_item"`
	DeleteAt        gorm.DeletedAt            `json:"-"`
}

func (AllTransactionResponse) TableName() string {
	return "transactions"
}

func (TransactionResponseId) TableName() string {
	return "transactions"
}

// {
//     "date": "2023-06-29T19:20:00Z",
//     "total": 3000000,
//     "status": "unpaid",
//     "user_id": 3
// }
