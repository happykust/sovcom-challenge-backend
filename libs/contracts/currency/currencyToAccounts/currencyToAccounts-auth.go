package currencyToAccounts

type ValidateRequest struct {
	AccessToken string `json:"access_token"`
}

type ValidateResponse struct {
	Status bool `json:"status"`
	UserID uint `json:"user_id"`
}
