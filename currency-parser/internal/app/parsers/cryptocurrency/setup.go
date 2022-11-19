package cryptocurrency

import (
	logger "currency-parser/pkg/logging"
	LoggerTypes "currency-parser/pkg/logging/types"
	"gonum.org/v1/gonum/stat/combin"
)

func Setup() {
	logger.Log(LoggerTypes.INFO, "[Currency-parser | Cryptocurrency] Start cryptocurrency parser", nil)

	TickersGroupID = GetGroupID(TickersGroupName)
	tickersSimpleCurrencyID := GetGroupID(TickersSimpleCurrencyGroupName)

	tickersCrypto := GetTickersInGroup(TickersGroupID)
	tickersSimple := GetTickersInGroup(tickersSimpleCurrencyID)

	for _, tickerCrypto := range tickersCrypto {
		for _, tickersSimple := range tickersSimple {
			go Consumer(tickerCrypto.TickerParse, tickersSimple.TickerParse)
		}
	}

	listIndexes := combin.Combinations(len(tickersCrypto), 2)
	for _, pairTickersIndexes := range listIndexes {
		go Consumer(tickersCrypto[pairTickersIndexes[1]].TickerParse, tickersCrypto[pairTickersIndexes[0]].TickerParse)
	}
}
