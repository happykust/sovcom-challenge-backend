package simpleDeals

import (
	"gorm.io/gorm"
)

type SimpleDeal struct {
	gorm.Model
	UserID      uint
	TickerGroup string
	TickerFrom  string
	TickerTo    string
	Amount      float64
	Currency    float64
}
