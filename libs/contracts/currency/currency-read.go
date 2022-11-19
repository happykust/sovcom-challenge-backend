package currency

type ReadRequest struct{}

type Ticker struct {
	Ticker string `json:"ticker"`
	Group  Group  `json:"group"`
}

type ReadResponse struct {
	Tickers []Ticker
}
