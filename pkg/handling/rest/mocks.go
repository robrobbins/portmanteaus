package rest

import (
	"github.com/robrobbins/portmanteaus/pkg/packing"
	"github.com/stretchr/testify/mock"
)

type PackingRecorderMock struct {
	mock.Mock
}

// Record allows the mock to implement Recorder
func (r *PackingRecorderMock) Record(s string, p *packing.Port) error {
	args := r.Called(s, p)
	return args.Error(0)
}
