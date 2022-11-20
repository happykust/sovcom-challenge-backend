package simpleDeals

import (
	"deals/internal/app/domain/simpleDeals/sending/amqp"
	"libs/contracts/currency"
	simpleDealsLib "libs/contracts/deals/simple"
	"libs/contracts/payments"
)

func SimpleDealBuyHandler(request simpleDealsLib.SimpleDealBuyRequest) simpleDealsLib.SimpleDealBuyResponse {
	reverseDeal := false

	tickerGroupCurrency := GetTickerCurrency(currency.ReadTickerRequest{TickerGroup: request.TickerGroup})
	if tickerGroupCurrency.Currency == 0 {
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Попробуйте позже. (1)"}
	}

	if request.TickerTo != tickerGroupCurrency.TickerTo && request.TickerTo != tickerGroupCurrency.TickerFrom {
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Неверный тикер для покупки."}
	}
	if request.TickerFrom != tickerGroupCurrency.TickerTo && request.TickerFrom != tickerGroupCurrency.TickerFrom {
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Неверный тикер для покупки."}
	}

	if request.TickerFrom == tickerGroupCurrency.TickerTo {
		reverseDeal = true
	}

	var cost float64
	cost = request.Amount / tickerGroupCurrency.Currency
	if reverseDeal {
		cost = request.Amount * tickerGroupCurrency.Currency
	}

	userTickerFromBalance := amqp.GetUserBalances(payments.GetBalancesRequest{UserID: request.UserID,
		Ticker: request.TickerFrom})

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
			Message: "Произошла ошибка при проведении сделки. Попробуйте позже. (2)"}
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
		SimpleDealCreateRepository(SimpleDeal{UserID: request.UserID, TickerGroup: request.TickerGroup,
			TickerFrom: request.TickerFrom, TickerTo: request.TickerTo, Amount: request.Amount, Currency: cost})
		return simpleDealsLib.SimpleDealBuyResponse{Status: false,
			Message: "Произошла ошибка при проведении сделки. Обратитесь в тех. поддержку."}
	}

	SimpleDealCreateRepository(SimpleDeal{UserID: request.UserID, TickerGroup: request.TickerGroup,
		TickerFrom: request.TickerFrom, TickerTo: request.TickerTo, Amount: request.Amount, Currency: cost})

	return simpleDealsLib.SimpleDealBuyResponse{
		Status: true, Message: "Сделка успешно проведена.", TickerGroup: request.TickerGroup,
		TickerFrom: request.TickerFrom, TickerTo: request.TickerTo,
		TickerFromBalance: updatedTickerFromBalance.Amount, TickerToBalance: updatedTickerToBalance.Amount,
		Amount: request.Amount}
}
