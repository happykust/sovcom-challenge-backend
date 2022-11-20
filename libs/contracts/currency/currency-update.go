package currency

type CurrencyUpdateRequestToCurrency struct {
	TickerGroup string      `json:"ticker_group"`
	TickerFrom  string      `json:"ticker_from"`
	TickerTo    string      `json:"ticker_to"`
	Data        interface{} `json:"data"`
}

type CurrencyUpdateRequestToDeals struct {
	TickerGroup string  `json:"ticker_group"`
	TickerFrom  string  `json:"ticker_from"`
	TickerTo    string  `json:"ticker_to"`
	Currency    float64 `json:"currency"`
}
