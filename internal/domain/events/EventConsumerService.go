package events

type EventConsumerService interface {
	ReadMessage() (*EventMessage, error)
	Close() error
}
