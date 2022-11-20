package promocodes

type DeleteRequest struct {
	Promocode string `json:"promocode"`
}

type DeleteResponse DeleteRequest
