package simpleDeals

type SimpleDealSellRequest struct {
	UserID      uint    `json:"user_id"`
	TickerGroup string  `json:"ticker_group"`
	TickerFrom  string  `json:"ticker_from"`
	TickerTo    string  `json:"ticker_to"`
	Amount      float64 `json:"amount"`
}

type SimpleDealSellResponse struct {
	Status            bool    `json:"status"`
	Message           string  `json:"message"`
	TickerGroup       string  `json:"ticker"`
	TickerFrom        string  `json:"ticker_from"`
	TickerTo          string  `json:"ticker_to"`
	TickerFromBalance float64 `json:"ticker_from_balance"`
	TickerToBalance   float64 `json:"ticker_to_balance"`
	Amount            float64 `json:"amount"`
}
