// Package message defines the structures and behavior of the framework messages.
package message

import (
	"fmt"
	"io"

	"github.com/google/uuid"
)

// Message is a container for events, commands, and queries.
type Message struct {
	// ID is a unique message correlation ID assigned by the framework.
	ID uuid.UUID

	// Data is the marshaled message payload.
	Data []byte
}

// New creates a new Message object from a data reader.
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
