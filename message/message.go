package message

import (
	"fmt"
	"io"

	"github.com/google/uuid"
)

type Message struct {
	ID   uuid.UUID
	Data []byte
}

func New(r io.Reader) (*Message, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error generating message ID: %w", err)
	}

	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("error reading message data: %w", err)
	}

	return &Message{ID: id, Data: buf}, nil
}
