package web_socket

import (
	dm2 "business/dm"
	"data/dm"
	"domain"
	"encoding/json"
	"fmt"
	"time"
)

func SendDmMessageHandler(event Event, c *Client) error {
	var chatevent SendDmMessageEvent

	if err := json.Unmarshal(event.Payload, &chatevent); err != nil {
		return fmt.Errorf("bad payload in request %v", err)
	}

	input := domain.DmMessage{
		SenderId: chatevent.SenderId,
		DmID:     chatevent.DmId,
		Message:  chatevent.Message,
		Username: chatevent.Username,
	}

	_, inputErr := dm.AddMessage(input)

	if inputErr != nil {
		return fmt.Errorf("Message was not saved %v", inputErr)
	}

	users, dmErr := dm2.FetchById(chatevent.DmId)

	if dmErr != nil {
		return fmt.Errorf("Message was not saved %v", inputErr)
	}

	var broadMessage NewDmMessageEvent

	broadMessage.Sent = time.Now()
	broadMessage.Message = chatevent.Message
	broadMessage.DmId = chatevent.DmId
	broadMessage.Username = chatevent.Username
	broadMessage.SenderId = chatevent.SenderId
	broadMessage.Id = chatevent.Id
	broadMessage.UsersId = users

	data, err := json.Marshal(broadMessage)

	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewDmMessage

	for client := range c.Manager.Clients {
		client.Egress <- outgoingEvent
	}

	return nil
}
