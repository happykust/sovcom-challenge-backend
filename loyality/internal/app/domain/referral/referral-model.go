package referral

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Referral struct {
	gorm.Model
	UUID          uuid.UUID `gorm:"type:uuid"`
	UserID        uint
	Ticker        string
	Amount        uint
	ReferralCount uint
}
