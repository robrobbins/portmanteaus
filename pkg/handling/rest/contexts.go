package rest

import (
	"github.com/julienschmidt/httprouter"
)

type ReadContext struct {
	Router *httprouter.Router
}

type RecordContext struct {
	Router   *httprouter.Router
	Recorder Recorder
}
