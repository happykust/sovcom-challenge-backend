package account

type AccountVerifyApprRequest struct {
	UserId    uint   `form:"user_id" json:"user_id" binding:"required"`
	RegStatus string `form:"reg_status" json:"reg_status" binding:"required"`
}

type AccountVerifyApprResponse struct {
	Yes bool
}
