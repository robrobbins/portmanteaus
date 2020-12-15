package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/julienschmidt/httprouter"
	"github.com/robrobbins/portmanteaus/pkg/handling/rest"
	"github.com/robrobbins/portmanteaus/pkg/storing/dynamo"
)

type app struct {
	router   *httprouter.Router
	dynamo   *dynamodb.DynamoDB // is safe for concurrent use...
	reader   *dynamo.ReadClient
	recorder *dynamo.RecordClient
}

func (a *app) setDynamo() {
	// establish an aws session, assume credential file...
	// TODO why won't aws find the region in a config file?
	s, err := session.NewSession(&aws.Config{Region: aws.String("us-west-1")})

	if err != nil {
		log.Fatalf("error creating aws session: %v", err)
	}

	// get a dynamo client instance and retain a reference on the app
	a.dynamo = dynamodb.New(s)
}

func (a *app) setReader() {}

func (a *app) setRecorder() {
	if a.dynamo == nil {
		log.Fatal("dynamo instance not set")
	}

	a.recorder = dynamo.NewRecordClient(a.dynamo)
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

func (a *app) setRecordingHandlers() {
	if a.recorder == nil || a.router == nil {
		log.Fatal("recorder and router instances are required")
	}

	// "Packing" as Humpty Dumpty said to Alice, "... there are two meanings packed up into one word."
	rest.AddPackingHandlers(&rest.RecordContext{
		Router:   a.router,
		Recorder: a.recorder,
	})
}
