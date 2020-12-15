package rest

import (
	"github.com/stretchr/testify/mock"
)

type RecorderMock struct {
	mock.Mock
}

// Record allows the mock to implement Recorder
func (r *RecorderMock) Record(a interface{}) error {
	args := r.Called(a)
	return args.Error(0)
}
