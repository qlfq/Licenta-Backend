package domain

type RoomPerson struct {
	Id     uint   `json:"id"`
	UserId uint   `json:"userId"`
	Label  string `json:"label"`
	RoomId uint   `json:"roomId"`
}
