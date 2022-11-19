package payments

type CreateBalanceUserRequest struct {
	UserID uint `json:"user_id"`
}

type CreateBalanceUserResponse struct {
	BalanceId uint `json:"balance_id"`
}
