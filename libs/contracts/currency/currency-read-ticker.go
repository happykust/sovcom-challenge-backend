package currency

type ReadTickerRequest struct {
	Ticker string `json:"ticker"`
}

type ReadTickerResponse struct {
	Ticker   string  `json:"ticker"`
	Currency float64 `json:"currency"`
}
