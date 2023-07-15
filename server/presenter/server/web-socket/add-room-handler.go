package web_socket

import (
	"business/rooms"
	"domain"
	"encoding/json"
	"fmt"
	"time"
)

func createEmptyPayload(c *Client, userId uint) {
	var broadMessage NewRoomEvent
	broadMessage.UserId = userId

	data, _ := json.Marshal(broadMessage)

	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewRoom

	for client := range c.Manager.Clients {
		client.Egress <- outgoingEvent
	}
}

func AddRoomHandler(event Event, c *Client) error {
	var roomEvent SendRoomEvent
	fmt.Println("here?")

	if err := json.Unmarshal(event.Payload, &roomEvent); err != nil {
		createEmptyPayload(c, 0)

		return nil
	}

	if roomEvent.Name == "" {
		createEmptyPayload(c, roomEvent.UserId)

		return nil
	}

	input := domain.Room{
		Name:    roomEvent.Name,
		UserId:  roomEvent.UserId,
		Persons: roomEvent.Persons,
		IsOpen:  roomEvent.IsOpen,
	}

	result, err := rooms.Add(input)

	if err != nil {
		fmt.Print("Here")
		createEmptyPayload(c, roomEvent.UserId)

		return nil
	}

	var broadMessage NewRoomEvent

	broadMessage.Sent = time.Now()
	broadMessage.Id = result.Id
	broadMessage.Name = result.Name
	broadMessage.Persons = result.Persons
	broadMessage.IsOpen = result.IsOpen

	data, err := json.Marshal(broadMessage)

	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewRoom

	for client := range c.Manager.Clients {
		client.Egress <- outgoingEvent
	}

	return nil
}
