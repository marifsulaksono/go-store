package entity

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	Id          int            `json:"id"`
	NameStore   string         `json:"name_store"`
	Address     string         `json:"address"`
	Email       string         `json:"email"`
	Description string         `json:"desc"`
	Status      string         `json:"status"`
	UserId      int            `json:"user_id"`
	CreateAt    time.Time      `json:"create_at"`
	DeleteAt    gorm.DeletedAt `json:"-"`
}

type StoreResponseById struct {
	Id           int       `json:"id"`
	NameStore    string    `json:"name_store"`
	Address      string    `json:"address"`
	Description  string    `json:"desc"`
	TotalProduct int       `json:"total_product"`
	Product      []Product `json:"product"`
	CreateAt     time.Time `json:"create_at"`
}

type ProductStoreResponse struct {
	Id        int    `json:"-"`
	NameStore string `json:"name_store"`
	Address   string `json:"address"`
}

func (ProductStoreResponse) TableName() string {
	return "stores"
}
