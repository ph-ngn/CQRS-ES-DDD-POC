package eventbus

import (
	"context"
	"encoding/json"
	"github.com/andyj29/wannabet/internal/domain"
	kafka "github.com/segmentio/kafka-go"
)

type EventBus struct {
	producer *kafka.Writer
}

func NewAsyncEventBus(producer *kafka.Writer) *EventBus {
	return &EventBus{
		producer: producer,
	}
}

func (b *EventBus) Publish(event domain.Event) {
	payload, err := json.Marshal(event)
	if err != nil {
		return
	}

	if err := b.producer.WriteMessages(context.Background(),
		kafka.Message{
			Topic: event.GetEventType(),
			Key:   []byte(event.GetAggregateID()),
			Value: payload,
		},
	); err != nil {
		// To implement storage fallback and retries
		return
	}
}
