package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int            `gorm:"primaryKey,autoIncrement" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Username    string         `gorm:"unique;not null;size:50" json:"username"`
	Password    string         `gorm:"not null" json:"password"`
	Email       string         `gorm:"unique;not null;size:255" json:"email"`
	Phonenumber *int           `gorm:"not null" json:"phonenumber"`
	Role        string         `gorm:"not null;default:buyer" json:"role"`
	CreateAt    time.Time      `gorm:"not null" json:"create_at"`
	UpdateAt    time.Time      `gorm:"default:null" json:"update_at"`
	DeleteAt    gorm.DeletedAt `json:"-"`
}

type UserResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Role        string `json:"role"`
}

type UserChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (UserResponse) TableName() string {
	return "users"
}

// {
// 	"name": "Muhammad Arif Sulaksono",
// 	"username": "arif",
// 	"password": "arif",
// 	"email": "arif@gmail.com",
// 	"phonenumber": 81234567890
// }
