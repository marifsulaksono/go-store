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
	Id           int               `json:"id"`
	NameStore    string            `json:"name_store"`
	Address      string            `json:"address"`
	Description  string            `json:"desc"`
	TotalProduct int               `json:"total_product"`
	Product      []ProductResponse `json:"product" gorm:"foreignKey:StoreId;references:Id"`
	CreateAt     time.Time         `json:"create_at"`
	DeleteAt     gorm.DeletedAt    `json:"-"`
}

type StoreResponse struct {
	Id        int    `json:"-"`
	NameStore string `json:"name_store"`
	Address   string `json:"address"`
	UserId    int    `json:"-"`
}

func (StoreResponseById) TableName() string {
	return "stores"
}

func (StoreResponse) TableName() string {
	return "stores"
}
