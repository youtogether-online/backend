package ws

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"sync"
	"time"
)

type Manager struct {
	clients  map[*client]bool
	upg      *websocket.Upgrader
	handlers map[string]EventHandler
	m        sync.RWMutex
	cfg      *conf.Config
}

func NewManager(cfg *conf.Config) *Manager {
	m := &Manager{clients: map[*client]bool{}, upg: &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}, handlers: map[string]EventHandler{}, cfg: cfg}

	m.handlers[eventMessage] = m.sendMessageHandler
	return m
}

func (m *Manager) Connect(c *gin.Context, roomId int) error {
	ws, err := m.upg.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return err
	}

	cl := m.newClient(ws, roomId, m.cfg.WS.PongWait)

	go m.readMessages(cl)
	go m.writeMessages(cl)
	return nil
}

func (m *Manager) routeEvent(event Event, c *client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("this event type is not supported")
	}
}

func (m *Manager) addClient(c *client) {
	if err := c.ws.SetReadDeadline(time.Now().Add(m.cfg.WS.PongWait)); err != nil {
		log.WithErr(err).Err("can't set read dead line")
	}

	c.ws.SetReadLimit(512)
	c.ws.SetPongHandler(c.pongHandler)

	m.m.Lock()
	defer m.m.Unlock()
	m.clients[c] = true
}

func (m *Manager) removeClient(c *client) {
	_ = c.ws.Close()
	m.m.Lock()
	defer m.m.Unlock()
	delete(m.clients, c)
}
