package web_socket

import (
	"business/messages"
	"business/rooms"
	"domain"
	"encoding/json"
	"fmt"
	"time"
)

func SendMessageHandler(event Event, c *Client) error {
	var chatevent SendMessageEvent

	if err := json.Unmarshal(event.Payload, &chatevent); err != nil {
		return fmt.Errorf("bad payload in request %v", err)
	}

	input := domain.Message{
		Username: chatevent.Username,
		UserId:   chatevent.SenderId,
		RoomId:   chatevent.RoomId,
		Text:     chatevent.Text,
	}

	_, inputErr := messages.Add(input)

	if inputErr != nil {
		return fmt.Errorf("Message was not saved %v", inputErr)
	}

	room := rooms.IsOpen(input.RoomId)

	usersId, usersErr := rooms.GetUsersByRoomId(input.RoomId)

	if usersErr != nil {
		return fmt.Errorf("failed to extract users %v", inputErr)
	}

	var broadMessage NewMessageEvent

	if room.Id != 0 {
		broadMessage.IsOpen = true
	}

	broadMessage.Sent = time.Now()
	broadMessage.Text = chatevent.Text
	broadMessage.SenderId = chatevent.SenderId
	broadMessage.RoomId = chatevent.RoomId
	broadMessage.Username = chatevent.Username
	broadMessage.UsersId = usersId

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewMessage

	for client := range c.Manager.Clients {
		client.Egress <- outgoingEvent
	}

	return nil
}
