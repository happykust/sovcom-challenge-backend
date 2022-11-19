package database

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Title string `json:"title"`
}

type Ticker struct {
	gorm.Model
	GroupID     uint   `json:"group_id"`
	Group       Group  `json:"group"`
	Ticker      string `gorm:"uniqueIndex" json:"ticker"`
	TickerParse string `json:"ticker_parse"`
}
