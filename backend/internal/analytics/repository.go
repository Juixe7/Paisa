package analytics

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetBudgetLimits(ctx context.Context, userID string) ([]*BudgetLimit, error)
	UpdateBudgetLimit(ctx context.Context, userID, categoryID string, limit float64) error
	GetMonthlyAggregates(ctx context.Context, userID, month string) ([]*MonthlyAggregate, error)
	
	GetSubscriptionDetections(ctx context.Context, userID string) ([]*SubscriptionDetection, error)
	CreateAnomalyLog(ctx context.Context, log *AnomalyLog) error
}

type pgRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgRepository{pool: pool}
}

func (r *pgRepository) GetBudgetLimits(ctx context.Context, userID string) ([]*BudgetLimit, error) {
	return nil, nil
}

func (r *pgRepository) UpdateBudgetLimit(ctx context.Context, userID, categoryID string, limit float64) error {
	return nil
}

func (r *pgRepository) GetMonthlyAggregates(ctx context.Context, userID, month string) ([]*MonthlyAggregate, error) {
	return nil, nil
}

func (r *pgRepository) GetSubscriptionDetections(ctx context.Context, userID string) ([]*SubscriptionDetection, error) {
	return nil, nil
}

func (r *pgRepository) CreateAnomalyLog(ctx context.Context, log *AnomalyLog) error {
	return nil
}
