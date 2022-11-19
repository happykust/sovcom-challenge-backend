package cryptocurrency

type CryptocurrencyKline struct {
	OpenPrice  any `json:"o"`
	ClosePrice any `json:"c"`
	HighPrice  any `json:"h"`
	LowPrice   any `json:"l"`
	BaseVolume any `json:"v"`
}

type CryptocurrencyIncoming struct {
	EventTime any                 `json:"E"`
	Kline     CryptocurrencyKline `json:"k"`
}
