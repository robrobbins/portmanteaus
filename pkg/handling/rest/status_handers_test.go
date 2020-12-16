package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type getStatusSuite struct {
	suite.Suite
	context *ReadContext
}

func (s *getStatusSuite) SetupSuite() {
	s.context = &ReadContext{
		Router: httprouter.New(),
	}

	// assure its only done 1x
	AddStatusHandlers(s.context)
}

func (s *getStatusSuite) TestGetHomeRoute() {
	r, _ := http.NewRequest("GET", HOME_ROUTE, nil)
	w := httptest.NewRecorder()
	s.context.Router.ServeHTTP(w, r)

	assert.Equal(s.T(), http.StatusOK, w.Code)
}

func (s *getStatusSuite) TestGetStatusRoute() {
	r, _ := http.NewRequest("GET", STATUS_ROUTE, nil)
	w := httptest.NewRecorder()
	s.context.Router.ServeHTTP(w, r)

	assert.Equal(s.T(), http.StatusOK, w.Code)
}

func TestGetStatusSuite(t *testing.T) {
	suite.Run(t, new(getStatusSuite))
}
