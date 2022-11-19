package promocodes

import "gorm.io/gorm"

type Promocode struct {
	gorm.Model
	Promocode            string `gorm:"uniqueIndex"`
	Ticker               string
	Amount               uint
	ActivationCountLimit uint
	ActivationCount      uint
}
