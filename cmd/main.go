package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	a := &app{}

	// get and assign the db abstraction here so that defer works as expected
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	defer db.Close()

	a.db = db
	a.setRouter()
	a.setStatusHandlers()
	a.setPackingHandlers()

	p := ":5000"

	fmt.Printf("Portmanteaus started, listening on %s.\n", p)

	if err := http.ListenAndServe(p, a.router); err != nil {
		log.Fatal("error staring app: ", err)
	}
}
