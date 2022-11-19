package referral

import "github.com/google/uuid"

type CreateRequest struct {
	UserID uint   `json:"user_id"`
	Ticker string `json:"ticker"`
	Amount uint   `json:"amount"`
}

type CreateResponse struct {
	UUID uuid.UUID `json:"uuid"`
}
