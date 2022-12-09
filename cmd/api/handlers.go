package main

import (
	"encoding/json"
	"fmt"
	"go-backend-template/internals/dtos"
	"go-backend-template/internals/dtos/requests"
	"io"
	"log"
	"net/http"
)

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := app.DB.AllMovies()

	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := json.Marshal(movies)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) GetBTCPrice(w http.ResponseWriter, r *http.Request) {

	// ------- #1 -----------------------------------------------------------------------
	response, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")

	if err != nil {
		log.Fatal(err)
	}

	bytes, errRead := io.ReadAll(response.Body)

	if errRead != nil {
		log.Fatal(errRead)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body) // Close when we exit close of this receiver function

	var currentPriceResponse requests.CurrentPriceResponse
	errUnmarshal := json.Unmarshal(bytes, &currentPriceResponse)

	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	// ----------------------------------------------------------------------------------

	// -------------------------- # 2 ---------------------------------------------------
	var coingeckoCurrencies, error = currentPriceResponse.BPI.GetCurrencies()

	if error != nil {
		log.Fatal(error)
	}

	// Find current fiat currencies
	fiatCurrencies, err := app.DB.GetFiatCurrencies()

	if fiatCurrencies == nil {
		log.Fatal("No fiat currency has been found")
	}
	// ------------------------------------------------------------------------------------

	// ------------- # 3 -------------------------------------------------------------------
	var res dtos.BtcSummaryDto

	// BPIS
	var bpis []requests.Currency

	for _, item := range coingeckoCurrencies {
		if _, ok := fiatCurrencies[item.Code]; ok {
			bpis = append(bpis, item)
		}
	}
	// --------------------------------------------------------------------------------------
	res.ChartName = currentPriceResponse.ChartName
	res.Disclaimer = currentPriceResponse.Disclaimer
	res.Time = currentPriceResponse.Time
	res.Bpi = dtos.Bpi{
		Currencies: bpis,
	}

	// final result
	out, err := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
