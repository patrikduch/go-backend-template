package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := app.DB.AllMovies()

	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := json.Marshal(movies)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
