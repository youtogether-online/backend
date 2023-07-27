package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"time"
)

type client struct {
	ws       *websocket.Conn
	event    chan Event
	pongWait time.Duration
	roomId   int
}

func (m *Manager) newClient(ws *websocket.Conn, roomId int, pongWait time.Duration) *client {
	c := &client{ws: ws, event: make(chan Event), roomId: roomId, pongWait: pongWait}
	m.addClient(c)
	return c
}

func (m *Manager) readMessages(c *client) {
	defer func() {
		m.removeClient(c)
	}()

	if err := c.ws.SetReadDeadline(time.Now().Add(c.pongWait)); err != nil {
		log.WithErr(err).Err("can't set read deadline")
		return
	}
	c.ws.SetReadLimit(512)
	c.ws.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.ws.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.WithErr(err).Err("unexpected error")
			}
			break
		}

		var req Event
		if err = json.Unmarshal(payload, &req); err != nil {
			log.WithErr(err).Err("can't unmarshal message")
			continue
		}

		if err = m.routeEvent(req, c); err != nil {
			log.WithErr(err).Err("can't handle message")
		}
	}
}

func (m *Manager) writeMessages(c *client) {
	ticker := time.NewTicker(m.cfg.WS.PingInterval)

	defer func() {
		ticker.Stop()
		// Graceful close if this triggers a closing
		m.removeClient(c)
	}()

	for {
		select {
		case message, ok := <-c.event:
			if !ok {
				if err := c.ws.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.WithErr(err).Err("connection closed", c.ws.RemoteAddr())
				}
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.WithErr(err).Err("can't marshal message")
				continue
			}

			if err = c.ws.WriteMessage(websocket.TextMessage, data); err != nil {
				log.WithErr(err).Err("can't write message")
				continue
			}
			log.Info("message was sent")
		case <-ticker.C:
			if err := c.ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.WithErr(err).Err("can't write message")
				return
			}
		}

	}
}

func (c *client) pongHandler(_ string) error {
	return c.ws.SetReadDeadline(time.Now().Add(c.pongWait))
}
