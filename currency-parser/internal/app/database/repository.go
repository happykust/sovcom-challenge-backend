package database

import "currency-parser/pkg/core/database"

func GetGroupIDRepository(groupTitle string) uint {
	var group database.Group
	database.PG.Where("title = ?", groupTitle).First(&group)
	return group.ID
}

func GetTickersInGroupRepository(groupID uint) []database.Ticker {
	var tickers []database.Ticker
	database.PG.Find(&tickers, "group_id = ?", groupID)
	return tickers
}
