package currency

type TickerDynamicJSONIncoming struct {
	Chart struct {
		Result []struct {
			Timestamp  []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Low    []float64 `json:"low"`
					High   []float64 `json:"high"`
					Open   []float64 `json:"open"`
					Volume []int64   `json:"volume"`
					Close  []float64 `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

type TickerDynamicKlineJSONOutcoming struct {
	Low    float64 `json:"l"`
	High   float64 `json:"h"`
	Open   float64 `json:"o"`
	Volume int64   `json:"v"`
	Close  float64 `json:"c"`
}

type TickerDynamicJSONOutcoming struct {
	Timestamp int64                           `json:"E"`
	Kline     TickerDynamicKlineJSONOutcoming `json:"k"`
}
