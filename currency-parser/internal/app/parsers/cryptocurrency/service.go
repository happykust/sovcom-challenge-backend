package cryptocurrency

import (
	"currency-parser/internal/app/database"
	database2 "currency-parser/pkg/core/database"
)

func GetGroupID() {
	TickersGroupID = database.GetGroupIDRepository(TickersGroupName)
}

func GetTickersInGroup() []database2.Ticker {
	var tickers []database2.Ticker
	tickers = database.GetTickersInGroupRepository(TickersGroupID)
	return tickers
}
