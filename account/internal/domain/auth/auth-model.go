package auth

import (
	"gorm.io/gorm"
)

type RegistrationStatus string

const (
	RegistrationStatusCreated  RegistrationStatus = "created"
	RegistrationStatusPending                     = "pending"
	RegistrationStatusDenied                      = "denied"
	RegistrationStatusApproved                    = "approved"
	RegistrationStatusVerified                    = "verified"
	RegisterStatusBlocked                         = "blocked"
)

type UnverifiedUsers struct {
	gorm.Model
	UserName           string
	Email              string
	FirstName          string
	LastName           string
	PasswordHash       string
	RefreshTokenHash   string
	RegistrationStatus RegistrationStatus `gorm:"default:'created'"`
	ReferralCode       string
	PersonalAssistant  uint
	MeetingInformation string
	AdditionalContact  string
}
