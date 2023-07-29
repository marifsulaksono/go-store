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
