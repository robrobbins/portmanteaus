package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/robrobbins/portmanteaus/pkg/handling/rest"
	"github.com/robrobbins/portmanteaus/pkg/recording/pgrecorder"
)

type app struct {
	router *httprouter.Router
	db     *sql.DB // long-lived database facade (not an actual connection)
}

func (a *app) setRouter() {
	a.router = httprouter.New()

	a.router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			rest.SetCommonHeaders(w, false)
			// TODO "Authorization will need to be added if we add auth...
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
}

func (a *app) setStatusHandlers() {
	if a.router == nil {
		log.Fatal("router instance is required")
	}

	rest.AddStatusHandlers(&rest.ReadContext{
		Router: a.router,
	})
}

func (a *app) setPackingHandlers() {
	if a.router == nil {
		log.Fatal("router instance is required")
	}

	// "Packing" as Humpty Dumpty said to Alice, "... there are two meanings packed up into one word."
	rest.AddPackingHandlers(&rest.PackingRecordContext{
		Router:   a.router,
		Recorder: pgrecorder.NewPackingRecorder(a.db), // AddPackingHandlers explicitly asks for a PackingRecorder via context
	})
}
