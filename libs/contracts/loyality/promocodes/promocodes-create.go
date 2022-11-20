package promocodes

type CreateRequest struct {
	Promocode            string `json:"promocode"`
	Ticker               string `json:"ticker"`
	Amount               uint   `json:"amount"`
	ActivationCountLimit uint   `json:"activation_count_limit"`
}

type CreateResponse CreateRequest
