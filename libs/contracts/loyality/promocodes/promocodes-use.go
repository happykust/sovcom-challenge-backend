package promocodes

type UseRequest struct {
	Promocode string `json:"promocode"`
}

type UseResponse struct {
	Message string `json:"message"`
}
