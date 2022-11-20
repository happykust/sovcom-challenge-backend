package referral

import "github.com/google/uuid"

type AddRequest struct {
	UUID       uuid.UUID `json:"uuid" required:"true"`
	ReferralID uint      `json:"referral_id"`
}

type AddResponse struct {
	Message string `json:"message"`
}
