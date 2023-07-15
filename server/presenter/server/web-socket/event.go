package web_socket

import (
	"domain"
	"encoding/json"
	"time"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

const (
	EventSendMessage   = "send_message"
	EventNewMessage    = "new_message"
	EventSendDmMessage = "send_dm_message"
	EventNewDmMessage  = "new_dm_message"
	EventSendNewRoom   = "send_room"
	EventNewRoom       = "add_new_room"
	EventDeleteRoom    = "delete_room"
	EventNewDeleteRoom = "room_deleted"
)

type SendMessageEvent struct {
	Text     string `json:"text"`
	SenderId uint   `json:"senderId"`
	RoomId   uint   `json:"roomId"`
	Username string `json:"username"`
	IsOpen   bool   `json:"isOpen"`
	UsersId  []uint `json:"usersId"`
}

type NewMessageEvent struct {
	SendMessageEvent
	Sent time.Time `json:"sent"`
}

type SendDmMessageEvent struct {
	Id       uint   `json:"id"`
	DmId     uint   `json:"dmId"`
	SenderId uint   `json:"senderId"`
	Username string `json:"username"`
	Message  string `json:"message"`
	UsersId  []uint `json:"usersId"`
}

type NewDmMessageEvent struct {
	SendDmMessageEvent
	Sent time.Time `json:"sent"`
}

type SendRoomEvent struct {
	Id      uint                `json:"id"`
	Name    string              `json:"name"`
	UserId  uint                `json:"userId"`
	Persons []domain.RoomPerson `json:"persons"`
	IsOpen  bool                `json:"isOpen"`
}

type NewRoomEvent struct {
	SendRoomEvent
	Sent time.Time `json:"sent"`
}

type SendDeleteRoom struct {
	RoomId  uint `json:"roomId"`
	Deleted bool `json:"deleted"`
}

type NewDeletedRoom struct {
	SendDeleteRoom
	Sent time.Time `json:"sent"`
}
