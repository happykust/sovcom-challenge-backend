package account

type AccountRefreshRequest struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" binding:"required"`
}

type AccountRefreshResponse struct {
	AccessToken  string `form:"access_token" json:"access_token"`
	RefreshToken string `form:"refresh_token" json:"refresh_token"`
}
