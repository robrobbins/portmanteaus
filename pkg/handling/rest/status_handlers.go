package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func status() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		respond(w, http.StatusOK, nil)
	}
}

func AddStatusHandlers(c *ReadContext) {
	c.Router.GET(HOME_ROUTE, status())
	c.Router.GET(STATUS_ROUTE, status())
}
