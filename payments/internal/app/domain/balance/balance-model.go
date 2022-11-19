package balance

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	BalanceId uint
	Ticker    string
	Amount    float64
}

type Balance struct {
	gorm.Model
	UserId uint
	Wallet []Wallet `gorm:"foreignKey:BalanceId"`
}
