package web_socket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Connection *websocket.Conn
	Manager    *Manager

	// egress is used to avoid concurrent writes on the websocket connection
	Egress chan Event
}

type List map[*Client]bool
