package categorisation

import "time"

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"` // expense, income
	CreatedAt time.Time `json:"created_at"`
}

type Subcategory struct {
	ID         string    `json:"id"`
	CategoryID string    `json:"category_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
}

type MerchantCache struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantKey   string    `json:"merchant_key"`
	CategoryID    string    `json:"category_id"`
	SubcategoryID *string   `json:"subcategory_id,omitempty"`
	Source        string    `json:"source"` // user_correction, ai_learned, rule
	Confidence    float64   `json:"confidence"`
	UseCount      int       `json:"use_count"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserCategoryCorrection struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	TransactionID string    `json:"transaction_id"`
	OldCategoryID string    `json:"old_category_id"`
	NewCategoryID string    `json:"new_category_id"`
	MerchantKey   string    `json:"merchant_key"`
	CreatedAt     time.Time `json:"created_at"`
}
