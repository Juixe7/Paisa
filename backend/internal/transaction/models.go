package transaction

import (
	"time"
)

type Direction string

const (
	Debit  Direction = "debit"
	Credit Direction = "credit"
)

type Source string

const (
	SmsSource    Source = "sms"
	ManualSource Source = "manual"
	PdfSource    Source = "pdf"
	IntentSource Source = "intent"
)

type Status string

const (
	Uncategorised Status = "uncategorised"
	Confirmed     Status = "confirmed"
	Flagged       Status = "flagged"
	Excluded      Status = "excluded"
)

type InflowType string

const (
	Salary        InflowType = "salary"
	Refund        InflowType = "refund"
	FriendPayback InflowType = "friend_payback"
	Cashback      InflowType = "cashback"
)

type Transaction struct {
	ID              string     `json:"id"`
	UserID          string     `json:"user_id"`
	Amount          float64    `json:"amount"`
	Direction       Direction  `json:"direction"`
	MerchantName    string     `json:"merchant_name"`
	VPA             *string    `json:"vpa,omitempty"`
	RawSMS          *string    `json:"raw_sms,omitempty"`
	Source          Source     `json:"source"`
	CategoryID      *string    `json:"category_id,omitempty"`
	SubcategoryID   *string    `json:"subcategory_id,omitempty"`
	ConfidenceScore *float64   `json:"confidence_score,omitempty"`
	Status          Status     `json:"status"`
	InflowType      *InflowType `json:"inflow_type,omitempty"`
	TransactedAt    time.Time  `json:"transacted_at"`
	CreatedAt       time.Time  `json:"created_at"`
}

type PendingIntent struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Amount       float64   `json:"amount"`
	MerchantName string    `json:"merchant_name"`
	CategoryID   string    `json:"category_id"`
	Status       string    `json:"status"` // pending, confirmed, expired
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type PDFUpload struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	S3Key     string    `json:"s3_key"`
	Status    string    `json:"status"` // pending, processing, completed, failed
	CreatedAt time.Time `json:"created_at"`
}
