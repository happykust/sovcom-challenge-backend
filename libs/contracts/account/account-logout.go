package account

type AccountLogoutRequest struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

type AccountLogoutResponse struct {
	Message string `form:"message" json:"message" binding:"required"`
}
