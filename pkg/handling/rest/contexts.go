package rest

import (
	"github.com/julienschmidt/httprouter"
)

type ReadContext struct {
	Router *httprouter.Router
}

type PackingRecordContext struct {
	Router   *httprouter.Router
	Recorder PackingRecorder
}
