package account

type AccountGetOneVUserRequest struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

type AccountGetOneVUserResponse struct {
	UserName     string `form:"user_name" json:"user_name" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required"`
	FirstName    string `form:"first_name" json:"first_name" binding:"required"`
	LastName     string `form:"last_name" json:"last_name" binding:"required"`
	Role         string `form:"role" json:"role" binding:"required"`
	ReferralCode string `form:"referral_code" json:"referral_code" binding:"required"`
	BalanceId    uint   `form:"balance_id" json:"balance_id" binding:"required"`
	ReferralId   uint   `form:"referral_id" json:"referral_id" binding:"required"`
	Ban          bool   `form:"ban" json:"ban" binding:"required"`
}
