package repositories

import (
	"database/sql"
	"go-backend-template/internals/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() (models.MovieSlice, error)
	GetFiatCurrencies() (map[string]*models.FiatCurrency, error)
}
