package account

type AccountSignUpRequest struct {
	Username     string `form:"username" json:"username" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required"`
	FirstName    string `form:"firstname" json:"firstname" binding:"required"`
	LastName     string `form:"lastname" json:"lastname" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	ReferralCode string `form:"referralcode" json:"referralCode"`
}

type AccountSignUpResponse struct {
	Message      string `form:"message" json:"message" binding:"required"`
	AccessToken  string `form:"access_token" json:"access_token"`
	RefreshToken string `form:"refresh_token" json:"refresh_token"`
}
