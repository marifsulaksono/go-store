package entity

import "time"

type Notification struct {
	Id                     int       `gorm:"primaryKey, autoIncrement" json:"id"`
	CreatedAt              time.Time `gorm:"not null" json:"created_at"`
	Title                  string    `gorm:"not null;size:255" json:"title"`
	Detail                 string    `gorm:"not null" json:"detail"`
	NotificationCategoryId int       `gorm:"not null" json:"notification_category_id"`
	RedirectUrl            string    `json:"redirect_url"`
	UserId                 int       `json:"user_id"`
}

type NotificationCategory struct {
	Id   int    `gorm:"primaryKey, autoIncrement" json:"id"`
	Name string `gorm:"not null;size:100" json:"name"`
}
