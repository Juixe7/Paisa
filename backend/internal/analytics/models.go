package analytics

import "time"

type BudgetLimit struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	CategoryID string    `json:"category_id"`
	Limit      float64   `json:"limit"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type MonthlyAggregate struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	CategoryID    string    `json:"category_id"`
	Month         string    `json:"month"` // YYYY-MM
	SpendTotal    float64   `json:"spend_total"`
	InflowTotal   float64   `json:"inflow_total"`
	UpdatedTimeAt time.Time `json:"updated_at"`
}

type SubscriptionDetection struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantName  string    `json:"merchant_name"`
	MonthlyAmount float64   `json:"monthly_amount"`
	FirstSeen     time.Time `json:"first_seen"`
	LastSeen      time.Time `json:"last_seen"`
}

type AnomalyLog struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	TransactionID string    `json:"transaction_id"`
	AnomalyType   string    `json:"anomaly_type"`
	Severity      string    `json:"severity"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
}
