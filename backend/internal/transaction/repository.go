package transaction

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateTransaction(ctx context.Context, tx *Transaction) error
	GetTransactionByID(ctx context.Context, id string) (*Transaction, error)
	ListTransactions(ctx context.Context, userID string, limit, offset int, filters map[string]interface{}) ([]*Transaction, error)
	CheckDuplicate(ctx context.Context, userID string, amount float64, direction Direction, transactedAt int64) (bool, error)
	
	CreatePendingIntent(ctx context.Context, intent *PendingIntent) error
	GetPendingIntentByID(ctx context.Context, id string) (*PendingIntent, error)
	UpdatePendingIntentStatus(ctx context.Context, id string, status string) error

	CreatePDFUpload(ctx context.Context, upload *PDFUpload) error
}

type pgRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgRepository{pool: pool}
}

func (r *pgRepository) CreateTransaction(ctx context.Context, tx *Transaction) error {
	return nil
}

func (r *pgRepository) GetTransactionByID(ctx context.Context, id string) (*Transaction, error) {
	return nil, nil
}

func (r *pgRepository) ListTransactions(ctx context.Context, userID string, limit, offset int, filters map[string]interface{}) ([]*Transaction, error) {
	return nil, nil
}

func (r *pgRepository) CheckDuplicate(ctx context.Context, userID string, amount float64, direction Direction, transactedAt int64) (bool, error) {
	return false, nil
}

func (r *pgRepository) CreatePendingIntent(ctx context.Context, intent *PendingIntent) error {
	return nil
}

func (r *pgRepository) GetPendingIntentByID(ctx context.Context, id string) (*PendingIntent, error) {
	return nil, nil
}

func (r *pgRepository) UpdatePendingIntentStatus(ctx context.Context, id string, status string) error {
	return nil
}

func (r *pgRepository) CreatePDFUpload(ctx context.Context, upload *PDFUpload) error {
	return nil
}
