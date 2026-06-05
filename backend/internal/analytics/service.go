package analytics

import (
	"context"
)

type SummaryResponse struct {
	TotalSpend      float64                `json:"total_spend"`
	TotalInflow     float64                `json:"total_inflow"`
	Velocity        string                 `json:"spending_velocity"` // on track, slightly over, etc.
	TopCategories   []map[string]interface{} `json:"top_categories"`
	InsightSentence string                 `json:"insight_sentence"`
}

type BudgetStatus struct {
	CategoryID string  `json:"category_id"`
	Limit      float64 `json:"limit"`
	Actual     float64 `json:"actual"`
}

type Service interface {
	GetSummary(ctx context.Context, userID string) (*SummaryResponse, error)
	GetBudgets(ctx context.Context, userID string) ([]*BudgetStatus, error)
	UpdateBudget(ctx context.Context, userID, categoryID string, limit float64) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetSummary(ctx context.Context, userID string) (*SummaryResponse, error) {
	// Mock implementation
	return &SummaryResponse{
		TotalSpend:      4500.00,
		TotalInflow:     15000.00,
		Velocity:        "on track",
		TopCategories:   []map[string]interface{}{
			{"category_id": "cat_groceries", "name": "Groceries", "amount": 1200.00},
			{"category_id": "cat_dining", "name": "Dining Out", "amount": 950.00},
			{"category_id": "cat_commute", "name": "Commute", "amount": 600.00},
		},
		InsightSentence: "You spent 12% less on Swiggy compared to last week. On track to save ₹3,000.",
	}, nil
}

func (s *service) GetBudgets(ctx context.Context, userID string) ([]*BudgetStatus, error) {
	// Mock implementation
	return []*BudgetStatus{
		{CategoryID: "cat_groceries", Limit: 5000.00, Actual: 1200.00},
		{CategoryID: "cat_dining", Limit: 3000.00, Actual: 950.00},
		{CategoryID: "cat_commute", Limit: 2000.00, Actual: 600.00},
	}, nil
}

func (s *service) UpdateBudget(ctx context.Context, userID, categoryID string, limit float64) error {
	return s.repo.UpdateBudgetLimit(ctx, userID, categoryID, limit)
}
