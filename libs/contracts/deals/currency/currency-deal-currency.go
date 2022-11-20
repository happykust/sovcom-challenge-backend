package currencyDeals

type IncomingCurrencyChangeRequest struct {
	TickerGroup string  `json:"ticker_group"`
	TickerFrom  string  `json:"ticker_from"`
	TickerTo    string  `json:"ticker_to"`
	Currency    float64 `json:"currency"`
}
