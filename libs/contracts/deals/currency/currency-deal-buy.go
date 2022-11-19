package currencyDeals

type CurrencyDealTrigger string

type CurrencyDealBuyRequest struct {
	UserID   uint                `json:"user_id"`
	Ticker   string              `json:"ticker"`
	Amount   float64             `json:"amount"`
	Currency float64             `json:"currency"`
	Trigger  CurrencyDealTrigger `json:"trigger"`
}

type CurrencyDealBuyResponse struct {
	Message string `json:"message"`
}
