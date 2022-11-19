package referral

import "github.com/google/uuid"

type UpdateRequest struct {
	UUID          uuid.UUID `json:"uuid" required:"true"`
	Ticker        string    `json:"ticker" required:"false"`
	Amount        uint      `json:"amount" required:"false"`
	ReferralCount uint      `json:"referral_count" required:"false"`
}

type UpdateResponse UpdateRequest
