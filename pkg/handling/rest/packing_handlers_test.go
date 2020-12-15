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

type handlingTestSuite struct {
	suite.Suite
	context *RecordContext // Client is the repository facade
}

func (s *handlingTestSuite) SetupSuite() {
	// NOTE if multiple tests existed in this suite we'd likely use SetupTest
	s.context = &RecordContext{
		Router:   httprouter.New(),
		Recorder: &RecorderMock{},
	}
}

func (s *handlingTestSuite) TestPostPortsRoute() {
	// setup the router
	AddPackingHandlers(s.context)
	// "spy" on the mocked repo facade. we do this to prevent the creation of integration tests.
	// this is a unit test that is only concerned with the returned http status
	// NOTE that you have to type assert the actual mock struct in order to assure the /mock package methods
	s.context.Recorder.(*RecorderMock).On("Record", mock.Anything).Return(nil)
	// a mocked http body for POSTing...
	body := getPostBody(`{"maker": "Bob", "first": "foo", "second": "bar"}`)
	r, _ := http.NewRequest("POST", PORTS_ROUTE, body)
	w := httptest.NewRecorder()
	s.context.Router.ServeHTTP(w, r)

	assert.Equal(s.T(), http.StatusOK, w.Code)

	// if you wanted to peek at what the mock was called with
	// args := s.context.Recorder.(*RecorderMock).Calls[0].Arguments
	// s.T().Log(args[0])
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(handlingTestSuite))
}
