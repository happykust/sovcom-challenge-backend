package simpleDeals

import (
	"deals/internal/app/domain/currencyDeals/types"
	"deals/internal/app/domain/simpleDeals/sending/amqp"
	"libs/contracts/currency"
	simpleDealsLib "libs/contracts/deals/simple"
	"libs/contracts/payments"
)

func SimpleDealBuyHandler(request simpleDealsLib.SimpleDealBuyRequest) simpleDealsLib.SimpleDealBuyResponse {
	userTIckerBalance := amqp.GetUserBalances(payments.GetBalancesRequest{UserID: request.UserID, Ticker: request.TickerFrom})

	tickerGroupCurrency := amqp.GetTickerCurrency(currency.ReadTickerRequest{TickerGroup: request.TickerGroup})
	if tickerGroupCurrency.Currency == 0 {
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Попробуйте позже."}
	}

	cost := request.Amount * ticker.Currency
	if userRUBBalance.Amount < cost {
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Недостаточно средств для проведения сделки."}
	}

	updatedRUBBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: request.UserID,
		Ticker: "RUB",
		Amount: -cost,
	})
	if updatedRUBBalance.Status == false {
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Попробуйте позже."}
	}

	updatedTickerBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: request.UserID,
		Ticker: request.Ticker,
		Amount: request.Amount,
	})

	if updatedTickerBalance.Status == false {
		amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
			UserID: request.UserID,
			Ticker: "RUB",
			Amount: cost,
		})
		SimpleDealCreateRepository(SimpleDeal{UserID: request.UserID, Type: types.BUY, Ticker: request.Ticker,
			Amount: request.Amount, Currency: ticker.Currency})
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Обратитесь в тех. поддержку."}
	}

	SimpleDealCreateRepository(SimpleDeal{UserID: request.UserID, Type: types.BUY, Ticker: request.Ticker,
		Amount: request.Amount, Currency: ticker.Currency})

	return simpleDealsLib.SimpleDealBuyResponse{RubBalance: updatedRUBBalance.Amount,
		Ticker: updatedTickerBalance.Ticker, Amount: updatedTickerBalance.Amount, Status: true,
		Message: "Сделка успешно проведена."}
}

func SimpleDealSellHandler(request simpleDealsLib.SimpleDealSellRequest) simpleDealsLib.SimpleDealSellResponse {
	userBalance := amqp.GetUserBalances(payments.GetBalancesRequest{UserID: request.UserID, Ticker: request.Ticker})

	ticker := amqp.GetTickerCurrency(currency.ReadTickerRequest{Ticker: request.Ticker})
	if ticker.Ticker == "" {
		return simpleDealsLib.SimpleDealSellResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Попробуйте позже."}
	}

	if userBalance.Amount < request.Amount {
		return simpleDealsLib.SimpleDealSellResponse{Status: false,
			Message: "Недостаточно средств для проведения сделки."}
	}

	cost := request.Amount * ticker.Currency

	updatedTickerBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: request.UserID,
		Ticker: request.Ticker,
		Amount: -request.Amount,
	})

	if updatedTickerBalance.Status == false {
		return simpleDealsLib.SimpleDealSellResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Попробуйте позже."}
	}

	updatedRUBBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: request.UserID,
		Ticker: "RUB",
		Amount: cost,
	})

	if updatedRUBBalance.Status == false {
		amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
			UserID: request.UserID,
			Ticker: request.Ticker,
			Amount: request.Amount,
		})
		SimpleDealCreateRepository(SimpleDeal{UserID: request.UserID, Type: types.SELL, Ticker: request.Ticker,
			Amount: request.Amount, Currency: ticker.Currency})
		return simpleDealsLib.SimpleDealSellResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Обратитесь в тех. поддержку."}
	}

	SimpleDealCreateRepository(SimpleDeal{UserID: request.UserID, Type: types.SELL, Ticker: request.Ticker,
		Amount: request.Amount, Currency: request.Currency})

	return simpleDealsLib.SimpleDealSellResponse{RubBalance: updatedRUBBalance.Amount,
		Ticker: updatedTickerBalance.Ticker, Amount: updatedTickerBalance.Amount, Status: true,
		Message: "Сделка успешно проведена."}
}
