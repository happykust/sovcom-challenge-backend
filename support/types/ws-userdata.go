package types

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type UserData struct {
	UserID    uint   `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      Role   `json:"role"`
}
