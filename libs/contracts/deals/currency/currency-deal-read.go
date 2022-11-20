package currencyDeals

type CurrencyDeal struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"UserID"`
	TickerGroup string  `json:"TickerGroup"`
	TickerFrom  string  `json:"TickerFrom"`
	TickerTo    string  `json:"TickerTo"`
	Amount      uint    `json:"Amount"`
	Currency    float64 `json:"Currency"`
	Trigger     string  `json:"Trigger"`
}

type CurrencyDealReadRequest struct {
	UserID uint `json:"user_id"`
}

type CurrencyDealReadResponse []CurrencyDeal
