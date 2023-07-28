package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	Stock      int            `json:"stock"`
	Price      int            `json:"price"`
	Sold       int            `json:"sold"`
	Desc       string         `json:"desc"`
	Status     string         `json:"status"`
	CategoryId int            `json:"category_id"`
	Category   Category       `json:"category"`
	DeleteAt   gorm.DeletedAt `json:"-"`
}

type ProductTransactionResponse struct {
	Id         int      `json:"-"`
	Name       string   `json:"name"`
	CategoryId int      `json:"category_id"`
	Category   Category `json:"category"`
}

func (ProductTransactionResponse) TableName() string {
	return "products"
}
