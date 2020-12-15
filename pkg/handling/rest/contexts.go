package rest

import (
	"github.com/julienschmidt/httprouter"
)

type RecordContext struct {
	Router   *httprouter.Router
	Recorder Recorder
}
