package currency

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	gorm.Model
	Title string `json:"title"`
}

type Ticker struct {
	gorm.Model
	GroupID uint   `json:"group_id"`
	Group   Group  `json:"group"`
	Ticker  string `gorm:"uniqueIndex"`
}

type TickerChange struct {
	gorm.Model
	Ticker   string    `json:"ticker"`
	Currency float32   `json:"currency"`
	Date     time.Time `json:"date"`
}
