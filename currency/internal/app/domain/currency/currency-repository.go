package currency

import (
	"currency/pkg/core/database"
)

func CurrencyReadRepository() []Ticker {
	var tickers []Ticker
	database.PG.Preload("Group").Find(&tickers)
	return tickers
}

func FindTickerRepositoryByTickerName(tickerName string) []Ticker {
	var tickers []Ticker
	database.PG.Preload("Group").First(&tickers, "ticker = ?", tickerName)
	return tickers
}

func CurrencyAddTickerRepository(ticker Ticker) Ticker {
	database.PG.Create(&ticker)
	return ticker
}

func CurrencyDeleteTickerRepository(ticker Ticker) string {
	database.PG.Delete(&ticker)
	return ticker.Ticker
}

func CurrencyCreateGroupRepository(group Group) Group {
	database.PG.Create(&group)
	return group
}

func FindGroupRepositoryByGroupTitle(groupTitle string) []Group {
	var groups []Group
	database.PG.First(&groups, "title = ?", groupTitle)
	return groups
}

func CurrencyReadGroupRepository() []Group {
	var groups []Group
	database.PG.Find(&groups, "id > ?", 0)
	return groups
}

func CurrencyDeleteGroupRepository(group Group) string {
	database.PG.Delete(&group)
	return group.Title
}

func CurrencyReadTickerChangesRepository(tickerName string) []TickerChange {
	var tickerChanges []TickerChange
	database.PG.Where("ticker = ?", tickerName).Find(&tickerChanges)
	return tickerChanges
}

func CurrencyReadTickerRepository(tickerTitle string) TickerChange {
	var ticker TickerChange
	database.PG.Where("ticker = ?", tickerTitle).Last(&ticker)
	return ticker
}
