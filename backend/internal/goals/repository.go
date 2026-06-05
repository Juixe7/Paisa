package goals

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateGoal(ctx context.Context, goal *Goal) error
	GetGoalByID(ctx context.Context, id string) (*Goal, error)
	ListGoals(ctx context.Context, userID string) ([]*Goal, error)
	UpdateGoalSavedAmount(ctx context.Context, id string, amount float64) error
	
	CreateContribution(ctx context.Context, contrib *GoalContribution) error
}

type pgRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgRepository{pool: pool}
}

func (r *pgRepository) CreateGoal(ctx context.Context, goal *Goal) error {
	return nil
}

func (r *pgRepository) GetGoalByID(ctx context.Context, id string) (*Goal, error) {
	return nil, nil
}

func (r *pgRepository) ListGoals(ctx context.Context, userID string) ([]*Goal, error) {
	return nil, nil
}

func (r *pgRepository) UpdateGoalSavedAmount(ctx context.Context, id string, amount float64) error {
	return nil
}

func (r *pgRepository) CreateContribution(ctx context.Context, contrib *GoalContribution) error {
	return nil
}
