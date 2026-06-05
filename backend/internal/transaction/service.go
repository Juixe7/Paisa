package transaction

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

var (
	ErrIntentExpired  = errors.New("payment intent has expired")
	ErrIntentNotFound = errors.New("payment intent not found")
)

type Service interface {
	CreateManualTransaction(ctx context.Context, userID string, amount float64, merchantName string, categoryID string) (*Transaction, error)
	IngestSMS(ctx context.Context, userID string, rawSMS string) (*Transaction, error)
	CreatePrePaymentIntent(ctx context.Context, userID string, amount float64, categoryID string, merchantName string) (*PendingIntent, error)
	ConfirmPrePaymentIntent(ctx context.Context, userID string, intentID string) (*Transaction, error)
}

type service struct {
	repo      Repository
	publisher Publisher
}

func NewService(repo Repository, publisher Publisher) Service {
	return &service{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *service) CreateManualTransaction(ctx context.Context, userID string, amount float64, merchantName string, categoryID string) (*Transaction, error) {
	// Generate UUID mock
	txID := "tx_" + time.Now().Format("20060102150405")
	tx := &Transaction{
		ID:           txID,
		UserID:       userID,
		Amount:       amount,
		Direction:    Debit,
		MerchantName: merchantName,
		Source:       ManualSource,
		CategoryID:   &categoryID,
		Status:       Confirmed,
		TransactedAt: time.Now(),
	}

	if err := s.repo.CreateTransaction(ctx, tx); err != nil {
		return nil, err
	}

	// Manual entries bypass AI pipeline but we can still publish transaction.created if analytics/goals need it.
	_ = s.publisher.PublishTransactionCreated(ctx, tx)

	return tx, nil
}

func (s *service) IngestSMS(ctx context.Context, userID string, rawSMS string) (*Transaction, error) {
	// Standard SMS hashing for deduplication
	hasher := sha256.New()
	hasher.Write([]byte(rawSMS))
	smsHash := hex.EncodeToString(hasher.Sum(nil))

	// In real setup, we verify if SMS is duplicate first.
	// For scaffolding, we mock transaction extraction
	amount := 150.00
	merchant := "Swiggy"
	direction := Debit

	isDup, err := s.repo.CheckDuplicate(ctx, userID, amount, direction, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	if isDup {
		return nil, errors.New("duplicate transaction detected")
	}

	txID := "tx_sms_" + smsHash[:10]
	tx := &Transaction{
		ID:           txID,
		UserID:       userID,
		Amount:       amount,
		Direction:    direction,
		MerchantName: merchant,
		RawSMS:       &rawSMS,
		Source:       SmsSource,
		Status:       Uncategorised,
		TransactedAt: time.Now(),
	}

	if err := s.repo.CreateTransaction(ctx, tx); err != nil {
		return nil, err
	}

	// Publish to RabbitMQ
	if err := s.publisher.PublishTransactionCreated(ctx, tx); err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *service) CreatePrePaymentIntent(ctx context.Context, userID string, amount float64, categoryID string, merchantName string) (*PendingIntent, error) {
	intentID := "intent_" + time.Now().Format("20060102150405")
	intent := &PendingIntent{
		ID:           intentID,
		UserID:       userID,
		Amount:       amount,
		MerchantName: merchantName,
		CategoryID:   categoryID,
		Status:       "pending",
		ExpiresAt:    time.Now().Add(5 * time.Minute),
	}

	if err := s.repo.CreatePendingIntent(ctx, intent); err != nil {
		return nil, err
	}

	return intent, nil
}

func (s *service) ConfirmPrePaymentIntent(ctx context.Context, userID string, intentID string) (*Transaction, error) {
	intent, err := s.repo.GetPendingIntentByID(ctx, intentID)
	if err != nil {
		return nil, err
	}
	if intent == nil {
		return nil, ErrIntentNotFound
	}
	if intent.Status != "pending" {
		return nil, errors.New("intent already resolved")
	}
	if time.Now().After(intent.ExpiresAt) {
		_ = s.repo.UpdatePendingIntentStatus(ctx, intentID, "expired")
		return nil, ErrIntentExpired
	}

	// Mark intent confirmed
	if err := s.repo.UpdatePendingIntentStatus(ctx, intentID, "confirmed"); err != nil {
		return nil, err
	}

	// Create transaction from intent
	txID := "tx_" + intentID[7:]
	tx := &Transaction{
		ID:           txID,
		UserID:       userID,
		Amount:       intent.Amount,
		Direction:    Debit,
		MerchantName: intent.MerchantName,
		Source:       IntentSource,
		CategoryID:   &intent.CategoryID,
		Status:       Confirmed,
		TransactedAt: time.Now(),
	}

	if err := s.repo.CreateTransaction(ctx, tx); err != nil {
		return nil, err
	}

	_ = s.publisher.PublishTransactionCreated(ctx, tx)

	return tx, nil
}
