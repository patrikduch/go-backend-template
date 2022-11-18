package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN string
	DB  *sql.DB
}

func main() {
	var app application

	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=movies"

	// connect to the database
	conn, err := app.connectTpDb()

	if err != nil {
		log.Fatal(err)
	}

	app.DB = conn

	defer app.DB.Close()

	log.Println("Starting application on port", port)

	// start a webserver
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
