package goals

import "time"

type Goal struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Title        string    `json:"title"`
	TargetAmount float64   `json:"target_amount"`
	SavedAmount  float64   `json:"saved_amount"`
	TargetDate   time.Time `json:"target_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GoalContribution struct {
	ID             string    `json:"id"`
	GoalID         string    `json:"goal_id"`
	Amount         float64   `json:"amount"`
	TransactionID  *string   `json:"transaction_id,omitempty"`
	RazorpayPaymentID *string `json:"razorpay_payment_id,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}

type SplitRequest struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"` // request initiator
	TransactionID string    `json:"transaction_id"`
	TotalAmount   float64   `json:"total_amount"`
	Status        string    `json:"status"` // pending, settled
	CreatedAt     time.Time `json:"created_at"`
}
