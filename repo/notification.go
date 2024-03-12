package repo

import (
	"context"
	"gostore/entity"

	"gorm.io/gorm"
)

type notificationRepository struct {
	DB *gorm.DB
}

type NotificationRepository interface {
	GetAllNotifications(ctx context.Context, userId int, categoryId int) ([]entity.Notification, error)
	InsertNotification(notif *entity.Notification) error
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{
		DB: db,
	}
}

func (n *notificationRepository) GetAllNotifications(ctx context.Context, userId int, categoryId int) ([]entity.Notification, error) {
	var (
		result []entity.Notification
		db     = n.DB
	)

	if categoryId > 0 {
		db.Where("notification_category_id = ?", categoryId)
	}

	err := db.Where("user_id = ? or user_id IS NULL", userId).Find(&result).Error
	return result, err
}

func (n *notificationRepository) InsertNotification(notif *entity.Notification) error {
	return n.DB.Create(&notif).Error
}
