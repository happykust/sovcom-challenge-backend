package types

type DealType string

type CurrencyDealTrigger string

const (
	BUY  DealType = "BUY"
	SELL          = "SELL"
)

const (
	UP   CurrencyDealTrigger = "UP"
	DOWN                     = "DOWN"
)
