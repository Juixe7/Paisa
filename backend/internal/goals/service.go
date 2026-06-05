package goals

import (
	"context"
	"time"
)

type Service interface {
	CreateGoal(ctx context.Context, userID, title string, targetAmount float64, targetDate time.Time) (*Goal, error)
	GetGoals(ctx context.Context, userID string) ([]*Goal, error)
	Contribute(ctx context.Context, goalID string, amount float64) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateGoal(ctx context.Context, userID, title string, targetAmount float64, targetDate time.Time) (*Goal, error) {
	goal := &Goal{
		ID:           "goal_" + time.Now().Format("20060102150405"),
		UserID:       userID,
		Title:        title,
		TargetAmount: targetAmount,
		SavedAmount:  0.0,
		TargetDate:   targetDate,
	}

	if err := s.repo.CreateGoal(ctx, goal); err != nil {
		return nil, err
	}

	return goal, nil
}

func (s *service) GetGoals(ctx context.Context, userID string) ([]*Goal, error) {
	return s.repo.ListGoals(ctx, userID)
}

func (s *service) Contribute(ctx context.Context, goalID string, amount float64) error {
	contrib := &GoalContribution{
		ID:        "contrib_" + time.Now().Format("20060102150405"),
		GoalID:    goalID,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
	if err := s.repo.CreateContribution(ctx, contrib); err != nil {
		return err
	}

	goal, err := s.repo.GetGoalByID(ctx, goalID)
	if err != nil {
		return err
	}
	if goal != nil {
		newSaved := goal.SavedAmount + amount
		return s.repo.UpdateGoalSavedAmount(ctx, goalID, newSaved)
	}

	return nil
}
