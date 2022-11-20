package promocodes

type ReadRequest struct {
	Promocode string `json:"promocode"`
}

type ReadResponse struct {
	Promocode            string `json:"promocode"`
	Ticker               string `json:"ticker"`
	Amount               uint   `json:"amount"`
	ActivationCountLimit uint   `json:"activation_count_limit"`
	ActivationCount      uint   `json:"activation_count"`
}
