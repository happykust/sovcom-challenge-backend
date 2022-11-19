package accounts

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type ValidateRequest struct {
	AccessToken string `json:"access_token"`
}

type ValidateResponse struct {
	Status    bool   `json:"status"`
	UserID    uint   `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      Role   `json:"role"`
}
