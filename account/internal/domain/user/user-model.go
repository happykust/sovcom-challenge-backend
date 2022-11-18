package user

import (
	"gorm.io/gorm"
)

type Role string

const (
	RoleDeveloper Role = "developer"
	RoleAdmin     Role = "admin"
	RoleUser      Role = "user"
)

type User struct {
	gorm.Model
	UserName         string
	Email            string
	FirstName        string
	LastName         string
	PasswordHash     string
	RefreshTokenHash string
	Role             Role `gorm:"default:'user'"`
	ReferralCode     string
	BalanceId        uint
	ReferralId       uint
}
