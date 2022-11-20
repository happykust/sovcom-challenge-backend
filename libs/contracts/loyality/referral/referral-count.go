package referral

type CountRequest struct {
	UserID uint `json:"user_id"`
}

type CountResponse struct {
	Count uint `json:"count"`
}
