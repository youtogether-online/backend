package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/internal/service"
	"net/http"
	"sync"
	"time"
)

var upg = websocket.Upgrader{
	HandshakeTimeout:  0,
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	WriteBufferPool:   nil,
	Subprotocols:      nil,
	Error:             nil,
	CheckOrigin:       nil,
	EnableCompression: false,
}

type Manager struct {
	clients  ClientList
	mutex    sync.Mutex
	handlers map[string]EventHandler
	otps     retentionMap
	redis    service.UserRedis
}

type ClientList map[*client]bool

func NewManager(ctx context.Context, redis service.UserRedis) *Manager {
	mg := &Manager{clients: make(ClientList), otps: newRetentionMap(ctx, 5*time.Second), handlers: make(map[string]EventHandler), redis: redis}
	mg.setupEventHandlers()
	return mg
}

func (m *Manager) setupEventHandlers() {
	m.handlers[eventMessage] = sendMessage
	m.handlers[eventChangeRoom] = chatRoomHandler
}

func chatRoomHandler(event Event, c *client) error {
	var changeRoomEvent ChangeRoomEvent
	if err := json.Unmarshal(event.Payload, &changeRoomEvent); err != nil {
		return err
	}
	c.chatRoom = changeRoomEvent.Name
	return nil
}

func (m *Manager) removeClient(c *client) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.clients[c]; ok {
		_ = c.conn.Close()
		delete(m.clients, c)
	}
}

func (m *Manager) addClient(c *client) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.clients[c] = true
}

func (m *Manager) routeEvent(event Event, c *client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("no such event type")
}

func (m *Manager) LoginHandler(c *gin.Context) error {

	var p dto.Password
	if err := json.NewDecoder(c.Request.Body).Decode(&p); err != nil {
		return err
	}

	//TODO auth
	if p.Password == "123" {
		OTP := m.otps.newOTP()

		resp := dao.OTP{OTP: OTP.key}
		c.JSON(http.StatusOK, resp)
		return nil
	}
	c.Status(http.StatusUnauthorized)
	return nil
}

func (m *Manager) Connect(w http.ResponseWriter, r *http.Request) error {
	conn, err := upg.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	cl := m.newClient(conn)

	go cl.readMessages()
	go cl.writeMessages()
	return nil
}

func (m *Manager) newClient(conn *websocket.Conn) *client {
	c := &client{conn: conn, m: m, egress: make(chan Event)}
	m.addClient(c)
	return c
}
