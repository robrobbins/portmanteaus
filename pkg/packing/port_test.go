package packing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortFactory(t *testing.T) {
	p := NewPort()

	if assert.NotNil(t, p, "it should return a Portmanteau") {
		assert.Equal(t, p.Algo, RANDO, "should be RANDO by default (for now)")
	}
}
