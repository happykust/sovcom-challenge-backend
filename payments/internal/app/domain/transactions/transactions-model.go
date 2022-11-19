package transactions

import (
	"gorm.io/gorm"
	"libs/contracts/payments"
)

type Transaction struct {
	gorm.Model
	TransactionUUID   string
	UserId            uint
	Amount            float64
	PromoCode         string                     `gorm:"default:'NO_PROMO'"`
	TransactionStatus payments.TransactionStatus `gorm:"default:PROCESSING"`
}
