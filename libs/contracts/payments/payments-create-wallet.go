package payments

type CreateWalletUserRequest struct {
	BalanceId uint    `json:"balance_id"`
	Ticker    string  `json:"ticker"`
	Amount    float64 `json:"ammount"`
}

type CreateWalletUserResponse struct {
	BalanceId uint `json:"balance_id"`
}
