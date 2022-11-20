package currencyDeals

type CurrencyDeal struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Type     string `json:"type"`
	Ticker   string `json:"ticker"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Trigger  string `json:"trigger"`
}

type CurrencyDealReadRequest struct {
	UserID uint `json:"user_id"`
}

type CurrencyDealReadResponse []CurrencyDeal
