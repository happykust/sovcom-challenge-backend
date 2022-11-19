package cryptocurrency

import (
	logger "currency-parser/pkg/logging"
	LoggerTypes "currency-parser/pkg/logging/types"
)

func Setup() {
	logger.Log(LoggerTypes.INFO, "[Currency-parser | Cryptocurrency] Start cryptocurrency parser", nil)
	GetGroupID()
	tickers := GetTickersInGroup()
	for _, ticker := range tickers {
		go Consumer(ticker.Ticker)
	}
}
