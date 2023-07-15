package domain

import (
	"time"
)

type Message struct {
	Id       uint      `json:"id"`
	Username string    `json:"username"`
	UserId   uint      `json:"userId"`
	RoomId   uint      `json:"roomId"`
	Date     time.Time `json:"date"`
	Text     string    `json:"text"`
}

const key = "jkadshSJKHFSD2432!34"

func (m *Message) Encrypt() {
	keyBytes := []byte(key)
	plaintextBytes := []byte(m.Text)

	for i := 0; i < len(plaintextBytes); i++ {
		plaintextBytes[i] ^= keyBytes[i%len(keyBytes)]
	}

	m.Text = string(plaintextBytes)
}

func (m *Message) Decrypt() {
	keyBytes := []byte(key)
	ciphertextBytes := []byte(m.Text)

	for i := 0; i < len(ciphertextBytes); i++ {
		ciphertextBytes[i] ^= keyBytes[i%len(keyBytes)]
	}

	m.Text = string(ciphertextBytes)
}
