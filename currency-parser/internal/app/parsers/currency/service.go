package currency

import (
	"currency-parser/internal/app/database"
	databaseCore "currency-parser/pkg/core/database"
)

func GetGroupID(groupName string) uint {
	return database.GetGroupIDRepository(groupName)
}

func GetTickersInGroup(groupID uint) []databaseCore.Ticker {
	var tickers []databaseCore.Ticker
	tickers = database.GetTickersInGroupRepository(groupID)
	return tickers
}
