package currency

type CurrencyUpdateRequest struct {
	TickerGroup string      `json:"ticker_group"`
	TickerFrom  string      `json:"ticker_from"`
	TickerTo    string      `json:"ticker_to"`
	Data        interface{} `json:"data"`
}
