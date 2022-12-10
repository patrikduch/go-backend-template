package dbrepo

import (
	"context"
	"database/sql"
	"go-backend-template/internals/models"
	"log"
	"time"
)

type PostgresDbRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDbRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDbRepo) GetFiatCurrencies() (map[string]*models.FiatCurrency, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	fiatCurrencyExecutor := models.FiatCurrencies()
	fiatCurrencies, err := fiatCurrencyExecutor.All(ctx, m.DB)

	if err != nil {
		log.Fatal(err)
	} // handle err

	if fiatCurrencies != nil {
		log.Println("aa")
	}

	mapResult := make(map[string]*models.FiatCurrency)

	for _, item := range fiatCurrencies {
		mapResult[item.CurrencyName] = item
	}

	return mapResult, nil
}

func (m *PostgresDbRepo) AllMovies() (models.MovieSlice, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	movieExecutor := models.Movies()
	movies, err := movieExecutor.All(ctx, m.DB)

	if err != nil {
		log.Fatal(err)
	} // handle err

	return movies, nil
}
