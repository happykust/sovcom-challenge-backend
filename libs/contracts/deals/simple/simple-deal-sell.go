package simpleDeals

type SimpleDealSellRequest struct {
	UserID   uint    `json:"user_id"`
	Ticker   string  `json:"ticker"`
	Amount   float64 `json:"amount"`
	Currency float64 `json:"currency"`
}

type SimpleDealSellResponse struct {
	Status     bool    `json:"status"`
	Message    string  `json:"message"`
	RubBalance float64 `json:"rub_balance"`
	Ticker     string  `json:"ticker"`
	Amount     float64 `json:"amount"`
}
