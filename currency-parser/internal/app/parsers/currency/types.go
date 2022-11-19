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

type TickerDynamicJSONOutcoming struct {
	Timestamp int64   `json:"timestamp"`
	Low       float64 `json:"low"`
	High      float64 `json:"high"`
	Open      float64 `json:"open"`
	Volume    int64   `json:"volume"`
	Close     float64 `json:"close"`
}
