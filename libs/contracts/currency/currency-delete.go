package currency

type DeleteRequest struct {
	Ticker string `json:"ticker"`
}

type DeleteResponse DeleteRequest
