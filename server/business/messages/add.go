package messages

import (
	"data/messages"
	"domain"
)

func Add(input domain.Message) (domain.Message, error) {
	input.Encrypt()

	result, err := messages.Add(input)

	if err != nil {
		return domain.Message{}, err
	}

	return result, nil
}
