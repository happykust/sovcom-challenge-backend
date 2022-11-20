package currencyDeals

import (
	"deals/internal/app/domain/currencyDeals/sending/amqp"
	"deals/internal/app/domain/currencyDeals/types"
	currencyDealsLib "libs/contracts/deals/currency"
	"libs/contracts/payments"
)

func CreateCurrencyBuyDealHandler(request currencyDealsLib.CurrencyDealBuyRequest) currencyDealsLib.CurrencyDealBuyResponse {
	CreateCurrencyDealRepository(CurrencyDeal{
		UserID:      request.UserID,
		TickerGroup: request.TickerGroup,
		TickerFrom:  request.TickerFrom,
		TickerTo:    request.TickerTo,
		Amount:      request.Amount,
		Currency:    request.Currency,
		Trigger:     types.CurrencyDealTrigger(request.Trigger),
	})
	return currencyDealsLib.CurrencyDealBuyResponse{Message: "OK"}
}

func ReadCurrencyDealHandler(request currencyDealsLib.CurrencyDealReadRequest) []CurrencyDeal {
	deals := GetCurrencyDealRepository(request.UserID)
	return deals
}

func DeleteCurrencyDealHandler(request currencyDealsLib.CurrencyDealDeleteRequest) currencyDealsLib.CurrencyDealDeleteResponse {
	DeleteCurrencyDealRepository(request.ID)
	return currencyDealsLib.CurrencyDealDeleteResponse{ID: request.ID}
}

func ExecuteCurrencyDealsHandler(request currencyDealsLib.IncomingCurrencyChangeRequest) {
	deals := GetAllCurrencyDealsByTickerRepository(request.TickerGroup)
	for _, deal := range deals {
		if (deal.Trigger == types.UP && request.Currency > deal.Currency) ||
			(deal.Trigger == types.DOWN && request.Currency < deal.Currency) {
			CurrencyDealBuyHandler(deal, request)
		}
	}
}

func CurrencyDealBuyHandler(deal CurrencyDeal, tickerChange currencyDealsLib.IncomingCurrencyChangeRequest) {
	reverseDeal := false

	if deal.TickerTo == tickerChange.TickerFrom {
		reverseDeal = true
	}

	userTickerFromBalance := amqp.GetUserBalances(payments.GetBalancesRequest{UserID: deal.UserID,
		Ticker: deal.TickerFrom})

	var cost float64
	cost = deal.Amount / tickerChange.Currency
	if reverseDeal {
		cost = deal.Amount * tickerChange.Currency
	}

	if userTickerFromBalance.Ticker == "" {
		ChangeCurrencyDealStatusRepository(deal.ID, true,
			"Произошла ошибка при попытке получения данных. Обратитесь в службу тех. поддержки.")
		return
	}

	if userTickerFromBalance.Amount < cost {
		ChangeCurrencyDealStatusRepository(deal.ID, true, "Недостаточно средств для проведения сделки.")
		return
	}

	updatedTickerFromBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: deal.UserID,
		Ticker: deal.TickerFrom,
		Amount: -cost,
	})
	if updatedTickerFromBalance.Status == false || updatedTickerFromBalance.Ticker == "" {
		ChangeCurrencyDealStatusRepository(deal.ID, false,
			"Невозможно обновить рублевой баланс пользователя.")
		return
	}

	updatedTickerToBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: deal.UserID,
		Ticker: deal.TickerTo,
		Amount: deal.Amount,
	})

	if updatedTickerToBalance.Status == false || updatedTickerToBalance.Ticker == "" {
		amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
			UserID: deal.UserID,
			Ticker: deal.TickerTo,
			Amount: cost,
		})
		ChangeCurrencyDealStatusRepository(deal.ID, true,
			"Невозможно осуществить ролл-бек баланса пользователя. Обратитесь в службу тех. поддержки.")
		return
	}

	ChangeCurrencyDealStatusRepository(deal.ID, true, "Сделка успешно проведена.")
}
