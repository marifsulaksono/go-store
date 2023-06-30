package entity

import "time"

type Transaction struct {
	Id              int                       `json:"id"`
	Date            time.Time                 `json:"date"`
	Total           int                       `json:"total"`
	Status          string                    `json:"status"`
	UserId          int                       `json:"user_id"`
	TransactionItem []TransactionItemResponse `gorm:"ForeignKey:TransactionId" json:"transaction_item"`
}

type TransactionResponse struct {
	Id              int                          `json:"id"`
	Date            time.Time                    `json:"date"`
	Total           int                          `json:"total"`
	Status          string                       `json:"status"`
	UserId          int                          `json:"user_id"`
	TransactionItem []AllTransactionItemResponse `gorm:"ForeignKey:TransactionId" json:"transaction_item"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
