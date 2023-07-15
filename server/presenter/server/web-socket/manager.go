package web_socket

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		Clients:  make(List),
		Handlers: make(map[string]EventHandler),
		Otps:     NewRetentionMap(ctx, 5*time.Second),
	}

	m.SetupEventHandlers()

	return m
}

func (m *Manager) ServeWS(c *gin.Context) {
	otp := c.Query("otp")

	if otp == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Please provide an OTP code")
		return
	}

	if !m.Otps.VerifyOTP(otp) {
		c.IndentedJSON(http.StatusUnauthorized, "Please provide an valid OTP code")
		return
	}

	conn, err := webSocketUpgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Print(err)

		return
	}

	newClient := NewClient(conn, m)

	m.addClient(newClient)

	// Start client processes (go routines)
	go newClient.readMessages()
	go newClient.writeMessages()
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	m.Clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Clients[client]; ok {
		client.Connection.Close()
		delete(m.Clients, client)
	}
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.Handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}

		return nil
	} else {
		return errors.New("there is no such event type")
	}
}

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	switch origin {
	case "http://localhost:3000":
		return true
	default:
		return true
	}
}
