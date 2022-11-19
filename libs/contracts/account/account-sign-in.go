package account

type AccountSignInRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type AccountSignInResponse struct {
	Message      string `json:"message"`
	AccessToken  string `form:"access_token" json:"access_token"`
	RefreshToken string `form:"refresh_token" json:"refresh_token"`
}
