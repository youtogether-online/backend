package ws

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	eventMessage    string = "send_message"
	eventNewMessage string = "new_message"
)

type Event struct {
	Type    string          `json:"type,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

type NewMessageEvent struct {
	SendMessageEvent
	Sent time.Time `json:"sent,omitempty"`
}

type SendMessageEvent struct {
	Msg  string `json:"msg,omitempty"`
	From string `json:"from,omitempty"`
}

type EventHandler func(event Event, c *client) error

func (m *Manager) sendMessageHandler(event Event, c *client) error {
	var chatEvent SendMessageEvent
	if err := json.Unmarshal(event.Payload, &chatEvent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	var broadMessage NewMessageEvent

	broadMessage.Sent = time.Now()
	broadMessage.Msg = chatEvent.Msg
	broadMessage.From = chatEvent.From

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	outgoingEvent := Event{
		Type:    eventNewMessage,
		Payload: data,
	}

	for cl := range m.clients {
		if cl.roomId == c.roomId {
			cl.event <- outgoingEvent
		}
	}
	return nil
}
