package main

import (
	"encoding/json"
	"fmt"
	"go-backend-template/internals/models"
	"net/http"
	"time"
)

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: rd,
		MPAARating:  "R",
		Runtime:     116,
		Description: "Very nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rd, _ = time.Parse("2006-01-02", "1981-06-12")

	rotla := models.Movie{
		ID:          2,
		Title:       "Raiders of the Lost Ark",
		ReleaseDate: rd,
		MPAARating:  "PG-13",
		Runtime:     115,
		Description: "Another very nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)
	movies = append(movies, rotla)

	out, err := json.Marshal(movies)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
