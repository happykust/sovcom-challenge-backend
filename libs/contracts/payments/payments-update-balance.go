package payments

type UpdateBalanceRequest struct {
	UserID uint    `json:"user_id"`
	Ticker string  `json:"ticker"`
	Amount float64 `json:"amount"`
}

type UpdateBalanceResponse struct {
	Status bool    `json:"status"`
	Ticker string  `json:"ticker"`
	Amount float64 `json:"amount"`
}
