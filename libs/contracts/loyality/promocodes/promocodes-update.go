package promocodes

type UpdateRequest struct {
	Promocode            string `json:"promocode" required:"true"`
	Ticker               string `json:"ticker" required:"false"`
	Amount               uint   `json:"amount" required:"false"`
	ActivationCountLimit uint   `json:"activation_count_limit" required:"false"`
}

type UpdateResponse struct {
	Promocode            string `json:"promocode"`
	Ticker               string `json:"ticker"`
	Amount               uint   `json:"amount"`
	ActivationCountLimit uint   `json:"activation_count_limit"`
}
