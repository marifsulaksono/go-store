package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Email       string         `json:"email"`
	Phonenumber *int           `json:"phonenumber"`
	Role        string         `json:"role"`
	CreateAt    time.Time      `json:"create_at"`
	UpdateAt    time.Time      `json:"update_at"`
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

type ShippingAddress struct {
	Id            int    `json:"id"`
	UserId        int    `json:"user_id"`
	RecipientName string `json:"recipient_name"`
	Address       string `json:"address"`
	Phonenumber   string `json:"phonenumber"`
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
