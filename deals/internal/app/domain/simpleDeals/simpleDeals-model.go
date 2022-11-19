package simpleDeals

import (
	"deals/internal/app/domain/currencyDeals/types"
	"gorm.io/gorm"
)

type SimpleDeal struct {
	gorm.Model
	UserID      uint
	Type        types.DealType
	TickerGroup string
	TickerFrom  string
	TickerTo    string
	Amount      float64
	Currency    float64
}
