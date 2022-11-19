package support

import "gorm.io/gorm"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type Ticket struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsOpen    bool   `json:"is_open" gorm:"default:true"`
}

type TicketMessage struct {
	gorm.Model
	TicketID uint   `json:"ticket_id"`
	Ticket   Ticket `json:"ticket"`
	FromID   uint   `json:"from_id"`
	Role     Role   `json:"role"`
	Message  string `json:"message"`
}
