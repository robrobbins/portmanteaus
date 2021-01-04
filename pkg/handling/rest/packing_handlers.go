package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/robrobbins/portmanteaus/pkg/packing" // NOTE left imports center as center dictates the contract
)

// packing is, hexagonally speaking, the "left" (application) that drives the "center" (domain). In this case, Packing.
// NOTE the RecordContext is passed from an "initializer". This is either cmd/main or test setup. All *_handlers will
// accept a context that is defined in rest/interfaces.
func pack(c *PackingRecordContext) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		p := packing.NewPort()

		if err := decodeBody(r, p); err != nil {
			// TODO logging
			respondHTTPErr(w, http.StatusBadRequest)
			return
		}

		// NOTE `p.Pack...` -> "Left drives center"
		if err := p.Pack(c.Recorder); err != nil {
			// TODO logging...
			respondErr(w, http.StatusInternalServerError, err)
			return
		}

		// NOTE not returning the newly constructed port, though we could...
		respond(w, http.StatusOK, nil)
	}
}

func AddPackingHandlers(c *PackingRecordContext) {
	c.Router.POST(PORTS_ROUTE, pack(c))
}
