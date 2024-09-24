package events

import (
	"log"
	"time"
)

type Event struct {
	Type      string      `json:"type"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

func NewEvent(eventType string, payload interface{}) Event {
	return Event{
		Type:      eventType,
		Payload:   payload,
		Timestamp: time.Now(),
	}
}

func PublishEvent(event Event) {
	log.Printf("Event published: Type=%s Payload=%s Timestamp=%v", event.Type, event.Payload, event.Timestamp)
}
