package referral

import "github.com/google/uuid"

type DeleteRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type DeleteResponse struct {
	UUID uuid.UUID `json:"uuid"`
}
