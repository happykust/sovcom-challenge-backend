package payments

type GetBalancesRequest struct {
	UserID uint   `json:"user_id"`
	Ticker string `json:"ticker"`
}

type GetBalancesResponse struct {
	Ticker string  `json:"ticker"`
	Amount float64 `json:"amount"`
}
