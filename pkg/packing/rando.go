package packing

import (
	"math/rand"
	"strings"
	"time"
)

//  randoPack is a simple algo which chooses random indices at which to splice the words together
func randoPack(p *Port) error {
	rand.Seed(time.Now().UnixNano())
	// get a len and index at which to stop for first...
	lenf := len(p.First)

	if lenf == 0 {
		// TODO logging on err
		return ErrFirstRequired
	}

	fi := rand.Intn(lenf - 1)
	// get len and rando for second...
	lens := len(p.Second)

	if lens == 0 {
		return ErrSecondRequired
	}

	si := rand.Intn(lens - 1)
	// ref the substrings...
	subf := p.First[:fi]
	subs := p.Second[si:lens]
	// build it
	var b strings.Builder
	b.Grow(lenf + lens)

	if _, err := b.WriteString(subf); err != nil {
		return err
	}

	if _, err := b.WriteString(subs); err != nil {
		return err
	}

	// append it
	p.Result = b.String()
	return nil
}
