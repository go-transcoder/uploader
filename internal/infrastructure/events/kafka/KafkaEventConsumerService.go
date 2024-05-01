package kafka

import (
	"context"
	"github.com/go-transcoder/uploader/internal/domain/events"
	"github.com/segmentio/kafka-go"
)

type EventConsumerService struct {
	address string
	reader  *kafka.Reader
}

func NewEventConsumerService(address string, destination string) *EventConsumerService {
	consumerService := &EventConsumerService{
		address: address,
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{address},
			GroupID:  "1231231231jnaja@consumer_group_test",
			Topic:    destination,
			MaxBytes: 10e6, // 10MB
		}),
	}

	return consumerService
}

func (e *EventConsumerService) ReadMessage() (*events.EventMessage, error) {
	m, err := e.reader.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}
	return events.NewEventMessage(m.Value, m.Offset), nil
}

func (e *EventConsumerService) Close() error {
	err := e.reader.Close()

	if err != nil {
		return err
	}
	return nil
}
