package packing

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version int    `json:"version"`
	Created int64  `json:"created"`
	Data    *Port  `json:"data"`
}

// NewEvent is a Factory producing an Event Aggregate
func NewEvent(d *Port) *Event {
	// we could consider that the invariant "all fields must be hydrated" is enforced here
	return &Event{
		ID:      uuid.New().String(),
		Name:    PORT_CREATED,
		Version: CURRENT_EVENT_VERSION,
		Created: time.Now().Unix(),
		Data:    d,
	}
}
