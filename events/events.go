package events

import (
	"log"
	"time"
)

type Event struct {
	Type      string    `json:"type"`
	Payload   []byte    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func (e *Event) String() string {
	return e.Type + ";" + string(e.Payload) + ";" + e.Timestamp.String()
}

func (h *Hub) NewEvent(eventType string, payload []byte) Event {
	return Event{
		Type:      eventType,
		Payload:   payload,
		Timestamp: time.Now(),
	}
}

// used for publishing to stdout
func (h *Hub) PublishEventOnLocal(event Event) {
	log.Printf("New Event: %s", event)
}
