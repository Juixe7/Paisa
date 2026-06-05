package categorisation

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	redisclient "github.com/redis/go-redis/v9"
)

type Service interface {
	CategoriseTransaction(ctx context.Context, userID, merchantName, vpa string) (string, string, float64, error)
	CorrectCategory(ctx context.Context, userID, txID, oldCategoryID, newCategoryID, subcategoryID, merchantName string) error
	GetCategories(ctx context.Context) ([]*Category, []*Subcategory, error)
}

type service struct {
	repo  Repository
	rdb   *redisclient.Client
}

func NewService(repo Repository, rdb *redisclient.Client) Service {
	return &service{
		repo: repo,
		rdb:  rdb,
	}
}

func (s *service) CategoriseTransaction(ctx context.Context, userID, merchantName, vpa string) (string, string, float64, error) {
	// Simple mock of the 3-layer engine:
	// Layer 1: Rule Engine
	// Layer 2: Personal Corrections Cache check
	// Layer 3: AI Fallback
	return "cat_groceries", "sub_vegetables", 0.98, nil
}

func (s *service) CorrectCategory(ctx context.Context, userID, txID, oldCategoryID, newCategoryID, subcategoryID, merchantName string) error {
	merchantKey := merchantName
	// 1. Save correction
	correction := &UserCategoryCorrection{
		ID:            "corr_" + time.Now().Format("20060102150405"),
		UserID:        userID,
		TransactionID: txID,
		OldCategoryID: oldCategoryID,
		NewCategoryID: newCategoryID,
		MerchantKey:   merchantKey,
		CreatedAt:     time.Now(),
	}
	
	if err := s.repo.CreateCorrection(ctx, correction); err != nil {
		return err
	}

	// 2. Update merchant cache mapping
	cache := &MerchantCache{
		ID:            "mcache_" + time.Now().Format("20060102150405"),
		UserID:        userID,
		MerchantKey:   merchantKey,
		CategoryID:    newCategoryID,
		Source:        "user_correction",
		Confidence:    1.0,
		UpdatedAt:     time.Now(),
	}
	if subcategoryID != "" {
		cache.SubcategoryID = &subcategoryID
	}
	
	if err := s.repo.UpsertMerchantCache(ctx, cache); err != nil {
		return err
	}

	// 3. Invalidate Redis Cache (as per Rule 7.1)
	merchantKeyHash := sha256Hash(merchantKey)
	catCacheKey := fmt.Sprintf("cat:%s:%s", userID, merchantKeyHash)
	
	currentMonth := time.Now().Format("2006-01")
	oldBudgetCacheKey := fmt.Sprintf("budget:%s:%s:%s", userID, currentMonth, oldCategoryID)
	newBudgetCacheKey := fmt.Sprintf("budget:%s:%s:%s", userID, currentMonth, newCategoryID)

	// MULTI/EXEC or Pipeline for atomic invalidation
	pipe := s.rdb.Pipeline()
	pipe.Del(ctx, catCacheKey)
	pipe.Del(ctx, oldBudgetCacheKey)
	pipe.Del(ctx, newBudgetCacheKey)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("invalidating redis cache: %w", err)
	}

	return nil
}

func (s *service) GetCategories(ctx context.Context) ([]*Category, []*Subcategory, error) {
	return s.repo.ListCategories(ctx)
}

func sha256Hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
