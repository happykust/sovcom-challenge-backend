package currencyDeals

type CurrencyDealTrigger string

type CurrencyDealBuyRequest struct {
	UserID      uint                `json:"user_id"`
	TickerGroup string              `json:"ticker_group"`
	TickerFrom  string              `json:"ticker_from"`
	TickerTo    string              `json:"ticker_to"`
	Amount      float64             `json:"amount"`
	Currency    float64             `json:"currency"`
	Trigger     CurrencyDealTrigger `json:"trigger"`
}

type CurrencyDealBuyResponse struct {
	Message string `json:"message"`
}
