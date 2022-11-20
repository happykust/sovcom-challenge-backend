package currency

type ReadTickerRequest struct {
	TickerGroup string `json:"ticker_group"`
}

type ReadTickerResponse struct {
	Currency   float64 `json:"currency"`
	TickerFrom string  `json:"tf"`
	TickerTo   string  `json:"tt"`
}
