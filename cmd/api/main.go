package main

import (
	"fmt"
	"go-backend-template/internals/repositories"
	"go-backend-template/internals/repositories/dbrepo"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN string
	DB  repositories.DatabaseRepo
}

func main() {
	var app application

	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=movies"

	// connect to the database
	conn, err := app.connectTpDb()

	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDbRepo{DB: conn}

	defer app.DB.Connection().Close()

	log.Println("Starting application on port", port)

	// start a webserver
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
