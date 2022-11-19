package currency

import (
	logger "currency-parser/pkg/logging"
	loggerTypes "currency-parser/pkg/logging/types"
	"gonum.org/v1/gonum/stat/combin"
)

func Setup() {
	logger.Log(loggerTypes.INFO, "[Currency-parser | Currency] Start currency parser", nil)

	TickersGroupID = GetGroupID(TickersGroupName)
	tickers := GetTickersInGroup(TickersGroupID)

	listIndexes := combin.Combinations(len(tickers), 2)
	for _, pairTickersIndexes := range listIndexes {
		go Consumer(tickers[pairTickersIndexes[0]].Ticker, tickers[pairTickersIndexes[1]].Ticker)
	}
}
