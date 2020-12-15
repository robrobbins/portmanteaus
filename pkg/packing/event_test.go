package packing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventFactory(t *testing.T) {
	// we'll require a portmanteau first
	d := NewPort()
	e := NewEvent(d)

	if assert.NotNil(t, e, "it should return an Event") {
		assert.Equal(t, e.Name, PORT_CREATED, "should be current version")
		assert.Equal(t, e.Version, CURRENT_EVENT_VERSION, "should be current version")
		assert.NotNil(t, e.Created)
		assert.Equal(t, e.Data, d)
	}
}
