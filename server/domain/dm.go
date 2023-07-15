package domain

type DM struct {
	Id           uint `json:"id"`
	FirstUserId  uint `json:"firstUserId"`
	SecondUserId uint `json:"secondUserId"`
}
