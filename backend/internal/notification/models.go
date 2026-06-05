package notification

import "time"

type Notification struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Type      string    `json:"type"` // info, warning, alert
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationLog struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	NotificationID string    `json:"notification_id"`
	Channel        string    `json:"channel"` // push, in_app
	Status         string    `json:"status"`  // sent, failed
	ErrorMessage   *string   `json:"error_message,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserNotificationPreferences struct {
	UserID             string    `json:"user_id"`
	EmailEnabled       bool      `json:"email_enabled"`
	PushEnabled        bool      `json:"push_enabled"`
	ThresholdPercent   int       `json:"threshold_percent"` // e.g., 80% or 100% budget
	UpdatedAt          time.Time `json:"updated_at"`
}
