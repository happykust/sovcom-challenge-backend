package currency

import (
	"libs/contracts/currency"
)

func CurrencyReadHandler() []Ticker {
	return CurrencyReadRepository()
}

func CurrencyAddTickerHandler(ticker currency.AddTickerRequest) currency.AddTickerResponse {
	isTickerExist := FindTickerRepositoryByTickerName(ticker.Ticker)
	if len(isTickerExist) != 0 {
		return currency.AddTickerResponse{}
	}
	isGroupExist := FindGroupRepositoryByGroupTitle(ticker.Group)
	if len(isGroupExist) == 0 {
		return currency.AddTickerResponse{}
	}
	CurrencyAddTickerRepository(Ticker{Ticker: ticker.Ticker, Group: isGroupExist[0]})
	return currency.AddTickerResponse{Ticker: ticker.Ticker, Group: currency.Group{Title: isGroupExist[0].Title, ID: isGroupExist[0].ID}}
}

func CurrencyDeleteTickerHandler(tickerName string) string {
	isTickerExist := FindTickerRepositoryByTickerName(tickerName)
	if len(isTickerExist) != 0 {
		return CurrencyDeleteTickerRepository(isTickerExist[0])
	}
	return "Ticker not found"
}

func CurrencyCreateTypeHandler(group Group) Group {
	isGroupExist := FindGroupRepositoryByGroupTitle(group.Title)
	if len(isGroupExist) == 0 {
		return CurrencyCreateGroupRepository(group)
	}
	return Group{}
}

func CurrencyReadTypeHandler() []Group {
	return CurrencyReadGroupRepository()
}

func CurrencyReadTickerChangesHandler(tickerName string) []TickerChange {
	return CurrencyReadTickerChangesRepository(tickerName)
}

func CurrencyDeleteTypeHandler(tickerTypeTitle string) string {
	isTickerTypeExist := FindGroupRepositoryByGroupTitle(tickerTypeTitle)
	if len(isTickerTypeExist) != 0 {
		return CurrencyDeleteGroupRepository(isTickerTypeExist[0])
	}
	return "Ticker type not found"
}

func CurrencyReadTickerHandler(request currency.ReadTickerRequest) currency.ReadTickerResponse {
	ticker := CurrencyReadTickerRepository(request.Ticker)
	return currency.ReadTickerResponse{Ticker: ticker.Ticker, Currency: float64(ticker.Currency)}
}
