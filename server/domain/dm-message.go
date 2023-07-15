package domain

import "time"

type DmMessage struct {
	Id       uint      `json:"id"`
	SenderId uint      `json:"senderId"`
	DmID     uint      `json:"dmId"`
	Message  string    `json:"message"`
	CreateAt time.Time `json:"createAt"`
	Username string    `json:"username"`
}
