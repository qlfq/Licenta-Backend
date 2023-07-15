package domain

import (
	"time"
)

type Room struct {
	Id        uint         `json:"id"`
	Name      string       `json:"name"`
	UserId    uint         `json:"userId"`
	Persons   []RoomPerson `json:"persons"`
	IsOpen    bool         `json:"isOpen"`
	createdAt time.Time
}
