package categorisation

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetCategoryByID(ctx context.Context, id string) (*Category, error)
	ListCategories(ctx context.Context) ([]*Category, []*Subcategory, error)
	
	GetMerchantCache(ctx context.Context, userID, merchantKey string) (*MerchantCache, error)
	UpsertMerchantCache(ctx context.Context, cache *MerchantCache) error
	
	CreateCorrection(ctx context.Context, correction *UserCategoryCorrection) error
}

type pgRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgRepository{pool: pool}
}

func (r *pgRepository) GetCategoryByID(ctx context.Context, id string) (*Category, error) {
	return nil, nil
}

func (r *pgRepository) ListCategories(ctx context.Context) ([]*Category, []*Subcategory, error) {
	return nil, nil, nil
}

func (r *pgRepository) GetMerchantCache(ctx context.Context, userID, merchantKey string) (*MerchantCache, error) {
	return nil, nil
}

func (r *pgRepository) UpsertMerchantCache(ctx context.Context, cache *MerchantCache) error {
	return nil
}

func (r *pgRepository) CreateCorrection(ctx context.Context, correction *UserCategoryCorrection) error {
	return nil
}
