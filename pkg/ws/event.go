package ws

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	eventMessage    string = "send message"
	eventNewMessage string = "new_message"
	eventChangeRoom string = "change_room"
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
	Msg  string `json:"message,omitempty"`
	From string `json:"from,omitempty"`
}

type ChangeRoomEvent struct {
	Name string `json:"name,omitempty"`
}

type EventHandler func(event Event, c *client) error

func sendMessage(event Event, c *client) error {
	var chatEvent SendMessageEvent
	if err := json.Unmarshal(event.Payload, &chatEvent); err != nil {
		return err
	}

	broadMessage := NewMessageEvent{
		SendMessageEvent: chatEvent,
		Sent:             time.Now(),
	}
	data, err := json.Marshal(broadMessage)
	if err != nil {
		return err
	}

	outGoingEvent := Event{
		Payload: data,
		Type:    eventNewMessage,
	}

	for cl := range c.m.clients {
		if cl.chatRoom == c.chatRoom {
			cl.egress <- outGoingEvent
		}
	}
	return nil
}

func checkOrigin(r *http.Request) bool {
	switch r.Header.Get("Origin") {
	case "http://localhost:3000":
		return true
	default:
		return false
	}
}
