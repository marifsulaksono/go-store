package entity

import (
	"gorm.io/gorm"
)

type Item struct {
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	Stock      int            `json:"stock"`
	Price      int            `json:"price"`
	IsSale     int            `json:"is_sale"`
	CategoryId int            `json:"category_id"`
	DeleteAt   gorm.DeletedAt `json:"delete_at"`
	Category   Category       `json:"category"`
	// Category   Category `gorm:"foreignKey:IdCategory" json:"category"` // inisialisasi foreignkey pada gorm
}

type ItemResponse struct {
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	Stock      int            `json:"stock"`
	Price      int            `json:"price"`
	IsSale     int            `json:"is_sale"`
	CategoryId int            `json:"category_id"`
	DeleteAt   gorm.DeletedAt `json:"-"`
	Category   Category       `json:"category"`
}

type ItemTransactionResponse struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
}

func (ItemResponse) TableName() string {
	return "items"
}

func (ItemTransactionResponse) TableName() string {
	return "items"
}
