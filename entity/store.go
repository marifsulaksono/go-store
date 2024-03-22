package entity

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	Id          int            `gorm:"primaryKey,autoIncrement" json:"id"`
	Name        string         `gorm:"not null;size:255" json:"name"`
	Address     string         `gorm:"not null" json:"address"`
	Email       string         `gorm:"not null" json:"email"`
	Description string         `json:"desc"`
	Status      string         `gorm:"not null;default:active" json:"status"`
	UserId      int            `gorm:"not null" json:"user_id"`
	CreateAt    time.Time      `gorm:"not null" json:"create_at"`
	DeleteAt    gorm.DeletedAt `json:"-"`
}

type StoreResponseById struct {
	Id           int               `json:"id"`
	Name         string            `json:"name"`
	Address      string            `json:"address"`
	Description  string            `json:"desc"`
	TotalProduct int               `json:"total_product"`
	Product      []ProductResponse `json:"product" gorm:"foreignKey:StoreId;references:Id"`
	CreateAt     time.Time         `json:"create_at"`
	DeleteAt     gorm.DeletedAt    `json:"-"`
}

type StoreResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	UserId  int    `json:"-"`
}

func (StoreResponseById) TableName() string {
	return "stores"
}

func (StoreResponse) TableName() string {
	return "stores"
}

// {
// 	"name_store": "Arif Cell",
// 	"address": "Kraksaan, Kabupaten Probolinggo",
// 	"email": "arif.cell@gmail.com",
// 	"desc": "Kami menyediakan berbagai macam kartu perdana"
// }
