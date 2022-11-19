package currency

type CurrencyUpdateRequest struct {
	TickerGroup string      `json:"ticker_group"`
	Data        interface{} `json:"data"`
}
