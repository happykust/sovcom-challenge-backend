package types

type DealType string

const (
	BUY  DealType = "BUY"
	SELL          = "SELL"
)

type GetCurrencyJSONResponse struct {
	EventTime int64 `json:"E"`
	Kline     struct {
		OpenPrice  float64 `json:"o"`
		ClosePrice float64 `json:"c"`
		HighPrice  float64 `json:"h"`
		LowPrice   float64 `json:"l"`
		BaseVolume int64   `json:"v"`
	} `json:"k"`
}
