package packing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRando(t *testing.T) {
	p := NewPort()

	p.First = "Musician"
	p.Second = "Engineer"

	assert.Nil(t, randoPack(p))

	t.Log(p)
}
