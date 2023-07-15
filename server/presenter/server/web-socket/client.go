package web_socket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		Connection: conn,
		Manager:    manager,
		Egress:     make(chan Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.Manager.removeClient(c)
	}()

	for {
		_, payload, err := c.Connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error reading: %v", err)
			}

			return
		}

		var request Event

		if err := json.Unmarshal(payload, &request); err != nil {
			log.Println("error marshalling event, ", err)

			break
		}

		if err := c.Manager.routeEvent(request, c); err != nil {
			log.Println("error handling message: ", err)
		}
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.Manager.removeClient(c)
	}()

	for {
		select {
		case message, ok := <-c.Egress:
			if !ok {
				if err := c.Connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("Connection closed: ", err)
				}

				return
			}

			fmt.Println("Is here")

			data, err := json.Marshal(message)

			if err != nil {
				log.Println(err)

				return
			}

			if err := c.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Printf("Failed to send message: %v", err)
			}

			log.Println("Message sent")
		}
	}
}
