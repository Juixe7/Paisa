package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type EventEnvelope struct {
	EventID   string      `json:"event_id"`
	EventType string      `json:"event_type"`
	Version   string      `json:"version"`
	Timestamp time.Time   `json:"timestamp"`
	UserID    string      `json:"user_id"`
	Payload   interface{} `json:"payload"`
}

type Publisher interface {
	PublishTransactionCreated(ctx context.Context, tx *Transaction) error
}

type rabbitPublisher struct {
	channel *amqp091.Channel
}

func NewPublisher(channel *amqp091.Channel) Publisher {
	return &rabbitPublisher{channel: channel}
}

func (p *rabbitPublisher) PublishTransactionCreated(ctx context.Context, tx *Transaction) error {
	payload := map[string]interface{}{
		"transaction_id": tx.ID,
		"amount":         tx.Amount,
		"direction":      tx.Direction,
		"merchant_name":  tx.MerchantName,
		"vpa":            tx.VPA,
		"source":         tx.Source,
		"transacted_at":  tx.TransactedAt,
	}

	envelope := EventEnvelope{
		EventID:   fmt.Sprintf("event-%s-%d", tx.ID, time.Now().UnixNano()),
		EventType: "transaction.created",
		Version:   "1",
		Timestamp: time.Now(),
		UserID:    tx.UserID,
		Payload:   payload,
	}

	body, err := json.Marshal(envelope)
	if err != nil {
		return fmt.Errorf("marshalling event envelope: %w", err)
	}

	err = p.channel.PublishWithContext(
		ctx,
		"transactions.exchange", // exchange
		"transaction.created",   // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("publishing message to rabbitmq: %w", err)
	}

	return nil
}
