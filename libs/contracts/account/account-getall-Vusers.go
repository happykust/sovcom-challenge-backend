package account

type AccountGettAllVUsersRequest struct {
	Get bool `form:"get" json:"get" binding:"required"`
}

type AccountGettAllVUsersResponse struct {
	Users []AccountGetOneVUserResponse `form:"users" json:"users" binding:"required"`
}
