package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	a := &app{}
	a.setDynamo()
	a.setRecorder()
	a.setRouter()
	a.setStatusHandlers()
	a.setRecordingHandlers()

	// aws eb default
	p := ":5000"

	fmt.Printf("Portmanteaus started, listening on %s.\n", p)

	if err := http.ListenAndServe(p, a.router); err != nil {
		log.Fatal("error staring app: ", err)
	}
}
