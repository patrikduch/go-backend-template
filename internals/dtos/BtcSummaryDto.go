package dtos

import "go-backend-template/internals/dtos/requests"

type Bpi struct {
	Currencies []requests.Currency `json:"currencies"`
}

type BtcSummaryDto struct {
	ChartName  string `json:"chartName"`
	Disclaimer string `json:"disclaimer"`
	Bpi        Bpi    `json:"bpi"`
	Time       struct {
		Updated string `json:"updated"`
	} `json:"time"`
}
