package events

import (
	"encoding/json"
)

type EventMessage struct {
	Offset  int64
	Message []byte
}

func NewEventMessage(message []byte, offset int64) *EventMessage {
	return &EventMessage{
		Offset:  offset,
		Message: message,
	}
}

func (e *EventMessage) GetObject() (interface{}, error) {
	var data interface{}
	err := json.Unmarshal(e.Message, &data)
	if err != nil {
		// Not a valid JSON-encoded string
		return nil, err
	}

	return data, nil
}
