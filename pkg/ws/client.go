package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"time"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type client struct {
	conn     *websocket.Conn
	m        *Manager
	egress   chan Event
	chatRoom string
}

func (c *client) readMessages() {
	defer c.m.removeClient(c)

	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.WithErr(err).Err("can't set read dead line")
	}

	c.conn.SetReadLimit(512)

	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.WithErr(err).Err("can't read websocket message")
			}
			break
		}

		var req Event
		if err = json.Unmarshal(payload, &req); err != nil {
			log.WithErr(err).Err("can't unmarshal event")
			break
		}

		if err = c.m.routeEvent(req, c); err != nil {
			log.WithErr(err).Err("can't handle message")
		}
	}
}

func (c *client) writeMessages() {
	defer c.m.removeClient(c)

	ticker := time.NewTicker(pingInterval)
	for {
		select {
		case msg, ok := <-c.egress:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.WithErr(err).Err("connection closed")
				}
				return
			}

			data, err := json.Marshal(msg)
			if err != nil {
				log.WithErr(err).Err("can't marshal message")
				return
			}

			if err = c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.WithErr(err).Err("failed to send message")
			}
			log.Info("message sent")
		case <-ticker.C:
			fmt.Println("ping")
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte("")); err != nil {
				log.WithErr(err).Err("failed to send ping message")
				return
			}
		}
	}
}

func (c *client) pongHandler(addr string) error {
	fmt.Println("pong")
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}
