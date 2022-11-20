package account

type AccountGetAllVUsersRequest struct {
	Get bool `form:"get" json:"get" binding:"required"`
}

type AccountGetAllVUsersResponse struct {
	Users []AccountGetOneVUserResponse `form:"users" json:"users" binding:"required"`
}

type AccountGetAllNVUsersResponse struct {
	Users []AccountGetOneNVUserResponse `form:"users" json:"users" binding:"required"`
}
