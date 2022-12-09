package repositories

import (
	"database/sql"
	"go-backend-template/internals/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
	GetFiatCurrencies() (map[string]*models.FiatCurrency, error)
}
