package dtos

import "go-backend-template/internals/dtos/requests"

type BtcItemDto struct {
	ChartName  string `json:"chartName"`
	Disclaimer string `json:"disclaimer"`
	Bpi        []requests.Currency
	Time       struct {
		Updated string `json:"updated"`
	} `json:"time"`
}
