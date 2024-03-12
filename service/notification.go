package service

import (
	"context"
	"gostore/entity"
	"gostore/repo"
	"time"
)

type notificationService struct {
	Repo repo.NotificationRepository
}

type NotificationService interface {
	GetAllNotifications(ctx context.Context, userId int, categoryId int) ([]entity.Notification, error)
	InsertNotification(notif *entity.Notification) error
}

func NewNotificationService(r repo.NotificationRepository) NotificationService {
	return &notificationService{
		Repo: r,
	}
}

func (n *notificationService) GetAllNotifications(ctx context.Context, userId int, categoryId int) ([]entity.Notification, error) {
	return n.Repo.GetAllNotifications(ctx, userId, categoryId)
}

func (n *notificationService) InsertNotification(notif *entity.Notification) error {
	notif.CreatedAt = time.Now()
	return n.Repo.InsertNotification(notif)
}
