package referral

import "github.com/google/uuid"

type ReadRequest struct {
	UUID   uuid.UUID `json:"uuid"`
	UserID uint      `json:"user_id"`
}

type ReadResponse struct {
	UserID        uint      `json:"user_id"`
	UUID          uuid.UUID `json:"uuid"`
	Ticker        string    `json:"ticker"`
	Amount        uint      `json:"amount"`
	ReferralCount uint      `json:"referral_count"`
	DepositBonus  uint      `json:"deposit_bonus"`
}
