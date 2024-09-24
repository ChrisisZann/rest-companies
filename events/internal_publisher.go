package events

type InternalPublisher struct {
	hub *Hub
}

func NewPublisher(h *Hub) *InternalPublisher {
	return &InternalPublisher{hub: h}
}

// used for publishing to websocket
func (c *InternalPublisher) WriteStream(event Event) {
	c.hub.broadcast <- []byte(event.String())
}
