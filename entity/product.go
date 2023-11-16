package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	Id         int            `gorm:"primaryKey,autoIncrement" json:"id"`
	Name       string         `gorm:"not null;size:255" json:"name"`
	Stock      *int           `gorm:"not null" json:"stock"`
	Price      *int           `gorm:"not null" json:"price"`
	Sold       int            `gorm:"null" json:"sold"`
	Desc       string         `json:"desc"`
	Status     string         `gorm:"not null" json:"status"`
	CategoryId *int           `gorm:"not null" json:"category_id"`
	Category   Category       `gorm:"-:migration" json:"category"`
	StoreId    int            `gorm:"not null" json:"store_id"`
	Store      StoreResponse  `gorm:"-:migration" json:"store"`
	DeleteAt   gorm.DeletedAt `json:"-"`
}

type ProductTransactionResponse struct {
	Id      int           `json:"-"`
	Name    string        `json:"name"`
	StoreId int           `json:"store_id"`
	Store   StoreResponse `json:"store"`
}

type ProductResponse struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Stock      int      `json:"stock"`
	Price      int      `json:"price"`
	Sold       int      `json:"sold"`
	StoreId    int      `json:"-"`
	CategoryId int      `json:"-"`
	Category   Category `json:"category"`
}

func (ProductTransactionResponse) TableName() string {
	return "products"
}

func (ProductResponse) TableName() string {
	return "products"
}
