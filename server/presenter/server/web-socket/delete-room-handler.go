package web_socket

import (
	"business/rooms"
	"encoding/json"
)

func createDeletedPayload(c *Client) {
	var broadMessage NewDeletedRoom
	broadMessage.Deleted = false

	data, _ := json.Marshal(broadMessage)

	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewDeleteRoom

	for client := range c.Manager.Clients {
		client.Egress <- outgoingEvent
	}
}

func DeleteRoomHandler(event Event, c *Client) error {
	var roomEvent SendDeleteRoom

	if err := json.Unmarshal(event.Payload, &roomEvent); err != nil {
		createDeletedPayload(c)

		return nil
	}

	result := rooms.Delete(roomEvent.RoomId)

	if result == false {
		createDeletedPayload(c)

		return nil
	}

	var broadMessage NewDeletedRoom
	broadMessage.Deleted = true

	data, _ := json.Marshal(broadMessage)

	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewDeleteRoom

	for client := range c.Manager.Clients {
		client.Egress <- outgoingEvent
	}

	return nil
}
