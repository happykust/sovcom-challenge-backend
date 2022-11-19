package currency

type ReadTickerRequest struct {
	TickerGroup string `json:"ticker_group"`
}

type ReadTickerResponse struct {
	Currency float64 `json:"currency"`
}
