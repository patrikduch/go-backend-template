package requests

type Currency struct {
	Code        string  `json:"code"`
	Symbol      string  `json:"symbol"`
	Rate        string  `json:"rate,omitempty"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}

type BPI struct {
	USD Currency `json:"usd"`
	GBP Currency `json:"gbp"`
	EUR Currency `json:"eur"`
}

type CurrentPriceResponse struct {
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	BPI        BPI    `json:"bpi"`

	Time struct {
		Updated string `json:"updated"`
	} `json:"time"`
}

func (u BPI) GetCurrencies() ([]Currency, error) {
	a := []Currency{
		u.GBP,
		u.EUR,
		u.USD,
	}
	return a, nil
}

//func (u BPI) MarshalJSONa() ([]byte, error) {
//	a := []interface{}{
//		u.GBP,
//		u.EUR,
//		u.USD,
//	}
//	return json.Marshal(a)
//}
