package currency

type AddTickerRequest struct {
	Group  string `json:"group"`
	Ticker string `json:"ticker"`
}

type AddTickerResponse struct {
	Group  Group  `json:"group"`
	Ticker string `json:"ticker"`
}
