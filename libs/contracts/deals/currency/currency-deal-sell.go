package currencyDeals

type CurrencyDealSellRequest struct {
	UserID   uint                `json:"user_id"`
	Ticker   string              `json:"ticker"`
	Amount   float64             `json:"amount"`
	Currency float64             `json:"currency"`
	Trigger  CurrencyDealTrigger `json:"trigger"`
}

type CurrencyDealSellResponse struct {
	Message string `json:"message"`
}
