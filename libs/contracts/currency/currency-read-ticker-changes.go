package currency

import "time"

type TickerChange struct {
	CreatedAt time.Time `json:"changes_at"`
	Ticker    string    `json:"ticker"`
	Currency  float32   `json:"currency"`
}

type ReadTickerChangesRequest struct {
	Ticker string `json:"ticker"`
}

type ReadTickerChangesResponse []TickerChange
