package currencyDeals

import (
	"deals/internal/app/domain/currencyDeals/types"
	"gorm.io/gorm"
)

type CurrencyDeal struct {
	gorm.Model
	UserID      uint
	TickerGroup string
	TickerFrom  string
	TickerTo    string
	Amount      float64
	Currency    float64
	Status      bool   `gorm:"default:false"`
	Tried       bool   `gorm:"default:false"`
	Message     string `gorm:"default:''"`
	Trigger     types.CurrencyDealTrigger
}
