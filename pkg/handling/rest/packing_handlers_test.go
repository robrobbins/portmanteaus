package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// NOTE create suites that are specific to a single, or a series of, tests
type postPortsSuite struct {
	suite.Suite
	context *PackingRecordContext
}

func (s *postPortsSuite) SetupSuite() {
	// NOTE SetupTest could be used as a `beforeEach` if needed...
	s.context = &PackingRecordContext{
		Router:   httprouter.New(),
		Recorder: &PackingRecorderMock{},
	}

	// setup the router
	AddPackingHandlers(s.context)
}

func (s *postPortsSuite) TestPostPortsRoute() {
	// "spy" on the mocked repo facade. we do this to prevent the creation of integration tests.
	// this is a unit test that is only concerned with the returned http status
	// NOTE that you have to type assert the actual mock struct in order to assure the /mock package methods
	s.context.Recorder.(*PackingRecorderMock).On("Record", mock.Anything, mock.Anything).Return(nil)
	// a mocked http body for POSTing...
	body := getPostBody(`{"maker": "Bob", "first": "foo", "second": "bar"}`)
	r, _ := http.NewRequest("POST", PORTS_ROUTE, body)
	w := httptest.NewRecorder()
	s.context.Router.ServeHTTP(w, r)

	assert.Equal(s.T(), http.StatusOK, w.Code)

	// if you wanted to peek at what the mock was called with
	// args := s.context.Recorder.(*PortRecorderMock).Calls[0].Arguments
	// s.T().Log(args[0])
}

func TestPostPortsSuite(t *testing.T) {
	suite.Run(t, new(postPortsSuite))
}
