package notification

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetNotifications(ctx context.Context, userID string) ([]*Notification, error)
	MarkAsRead(ctx context.Context, userID, notificationID string) error
	CreateNotification(ctx context.Context, notif *Notification) error
	LogNotification(ctx context.Context, log *NotificationLog) error
}

type pgRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgRepository{pool: pool}
}

func (r *pgRepository) GetNotifications(ctx context.Context, userID string) ([]*Notification, error) {
	return nil, nil
}

func (r *pgRepository) MarkAsRead(ctx context.Context, userID, notificationID string) error {
	return nil
}

func (r *pgRepository) CreateNotification(ctx context.Context, notif *Notification) error {
	return nil
}

func (r *pgRepository) LogNotification(ctx context.Context, log *NotificationLog) error {
	return nil
}
