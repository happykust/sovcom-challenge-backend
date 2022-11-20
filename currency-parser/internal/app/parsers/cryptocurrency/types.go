package cryptocurrency

type CryptocurrencyKline struct {
	OpenPrice  any    `json:"o"`
	ClosePrice string `json:"c"`
	HighPrice  any    `json:"h"`
	LowPrice   any    `json:"l"`
	BaseVolume any    `json:"v"`
}

type CryptocurrencyIncoming struct {
	EventTime  any                 `json:"E"`
	TickerFrom string              `json:"tf"`
	TickerTo   string              `json:"tt"`
	Kline      CryptocurrencyKline `json:"k"`
}
