package simpleDeals

import (
	"deals/internal/app/domain/currencyDeals/types"
	"deals/internal/app/domain/simpleDeals/sending/amqp"
	"libs/contracts/currency"
	simpleDealsLib "libs/contracts/deals/simple"
	"libs/contracts/payments"
)

func SimpleDealBuyHandler(request simpleDealsLib.SimpleDealBuyRequest) simpleDealsLib.SimpleDealBuyResponse {
	userTickerFromBalance := amqp.GetUserBalances(payments.GetBalancesRequest{UserID: request.UserID,
		Ticker: request.TickerFrom})

	tickerGroupCurrency := GetTickerCurrency(currency.ReadTickerRequest{TickerGroup: request.TickerGroup})
	if tickerGroupCurrency.Currency == 0 {
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Попробуйте позже."}
	}

	cost := request.Amount * tickerGroupCurrency.Currency
	if userTickerFromBalance.Amount < cost {
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Недостаточно средств для проведения сделки."}
	}

	updatedTickerFromBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: request.UserID,
		Ticker: request.TickerFrom,
		Amount: -cost,
	})
	if updatedTickerFromBalance.Status == false {
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Попробуйте позже."}
	}

	updatedTickerToBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: request.UserID,
		Ticker: request.TickerTo,
		Amount: request.Amount,
	})

	if updatedTickerToBalance.Status == false {
		amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
			UserID: request.UserID,
			Ticker: request.TickerFrom,
			Amount: cost,
		})
		SimpleDealCreateRepository(SimpleDeal{UserID: request.UserID, Type: types.BUY, TickerGroup: request.TickerGroup,
			TickerFrom: request.TickerFrom, TickerTo: request.TickerTo, Amount: request.Amount, Currency: cost})
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Обратитесь в тех. поддержку."}
	}

	SimpleDealCreateRepository(SimpleDeal{UserID: request.UserID, Type: types.BUY, TickerGroup: request.TickerGroup,
		TickerFrom: request.TickerFrom, TickerTo: request.TickerTo, Amount: request.Amount, Currency: cost})

	return simpleDealsLib.SimpleDealBuyResponse{
		Status: true, Message: "Сделка успешно проведена.", TickerGroup: request.TickerGroup,
		TickerFrom: request.TickerFrom, TickerTo: request.TickerTo,
		TickerFromBalance: updatedTickerFromBalance.Amount, TickerToBalance: updatedTickerToBalance.Amount,
		Amount: request.Amount}
}

func SimpleDealSellHandler(request simpleDealsLib.SimpleDealSellRequest) simpleDealsLib.SimpleDealSellResponse {
	userTickerFromBalance := amqp.GetUserBalances(payments.GetBalancesRequest{UserID: request.UserID,
		Ticker: request.TickerFrom})

	tickerGroupCurrency := GetTickerCurrency(currency.ReadTickerRequest{TickerGroup: request.TickerGroup})
	if tickerGroupCurrency.Currency == 0 {
		return simpleDealsLib.SimpleDealSellResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Попробуйте позже."}
	}

	if userTickerFromBalance.Amount < request.Amount {
		return simpleDealsLib.SimpleDealSellResponse{Status: false,
			Message: "Недостаточно средств для проведения сделки."}
	}

	cost := request.Amount * tickerGroupCurrency.Currency

	updatedTickerFromBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: request.UserID,
		Ticker: request.TickerFrom,
		Amount: -request.Amount,
	})

	if updatedTickerFromBalance.Status == false {
		return simpleDealsLib.SimpleDealSellResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Попробуйте позже."}
	}

	updatedTickerToBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: request.UserID,
		Ticker: request.TickerTo,
		Amount: cost,
	})

	if updatedTickerToBalance.Status == false {
		amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
			UserID: request.UserID,
			Ticker: request.TickerFrom,
			Amount: request.Amount,
		})

		SimpleDealCreateRepository(SimpleDeal{UserID: request.UserID, Type: types.SELL, TickerGroup: request.TickerGroup,
			Amount: request.Amount, Currency: tickerGroupCurrency.Currency, TickerFrom: request.TickerFrom,
			TickerTo: request.TickerTo})

		return simpleDealsLib.SimpleDealSellResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Обратитесь в тех. поддержку."}
	}

	SimpleDealCreateRepository(SimpleDeal{UserID: request.UserID, Type: types.SELL, TickerGroup: request.TickerGroup,
		Amount: request.Amount, Currency: tickerGroupCurrency.Currency, TickerFrom: request.TickerFrom,
		TickerTo: request.TickerTo})

	return simpleDealsLib.SimpleDealSellResponse{
		Status: true, Message: "Сделка успешно проведена.", TickerGroup: request.TickerGroup,
		TickerFrom: request.TickerFrom, TickerTo: request.TickerTo,
		TickerFromBalance: updatedTickerFromBalance.Amount, TickerToBalance: updatedTickerToBalance.Amount,
		Amount: request.Amount}
}
