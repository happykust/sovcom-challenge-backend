package types

type GetCurrencyJSONResponse struct {
	EventTime  int64  `json:"E"`
	TickerFrom string `json:"tf"`
	TickerTo   string `json:"tt"`
	Kline      struct {
		OpenPrice  string `json:"o"`
		ClosePrice string `json:"c"`
		HighPrice  string `json:"h"`
		LowPrice   string `json:"l"`
		BaseVolume string `json:"v"`
	} `json:"k"`
}
