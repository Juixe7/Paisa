package notification

import (
	"context"
	"time"
)

type Service interface {
	GetNotifications(ctx context.Context, userID string) ([]*Notification, error)
	MarkRead(ctx context.Context, userID, notificationID string) error
	SendPushNotification(ctx context.Context, userID, title, body string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetNotifications(ctx context.Context, userID string) ([]*Notification, error) {
	// Mock implementation
	return []*Notification{
		{
			ID:        "notif_1",
			UserID:    userID,
			Title:     "Budget Alert",
			Body:      "You have spent 82% of your Groceries budget.",
			Type:      "warning",
			Read:      false,
			CreatedAt: time.Now().Add(-10 * time.Minute),
		},
	}, nil
}

func (s *service) MarkRead(ctx context.Context, userID, notificationID string) error {
	return s.repo.MarkAsRead(ctx, userID, notificationID)
}

func (s *service) SendPushNotification(ctx context.Context, userID, title, body string) error {
	notif := &Notification{
		ID:        "notif_push_" + time.Now().Format("20060102150405"),
		UserID:    userID,
		Title:     title,
		Body:      body,
		Type:      "info",
		CreatedAt: time.Now(),
	}
	if err := s.repo.CreateNotification(ctx, notif); err != nil {
		return err
	}
	
	// FCM integration would go here.
	log := &NotificationLog{
		ID:             "notiflog_" + time.Now().Format("20060102150405"),
		UserID:         userID,
		NotificationID: notif.ID,
		Channel:        "push",
		Status:         "sent",
		CreatedAt:      time.Now(),
	}
	return s.repo.LogNotification(ctx, log)
}
