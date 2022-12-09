package requests

type Currency struct {
	Code        string  `json:"code"`
	Symbol      string  `json:"symbol"`
	Rate        string  `json:"rate,omitempty"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}

type BtcBpiResponse struct {
	USD Currency `json:"usd"`
	GBP Currency `json:"gbp"`
	EUR Currency `json:"eur"`
}

type CurrentPriceResponse struct {
	Disclaimer string         `json:"disclaimer"`
	ChartName  string         `json:"chartName"`
	BPI        BtcBpiResponse `json:"bpi"`

	Time struct {
		Updated string `json:"updated"`
	} `json:"time"`
}

func (u BtcBpiResponse) GetCurrencies() ([]Currency, error) {
	a := []Currency{
		u.GBP,
		u.EUR,
		u.USD,
	}
	return a, nil
}
