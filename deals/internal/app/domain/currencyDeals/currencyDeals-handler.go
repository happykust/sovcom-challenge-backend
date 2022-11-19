package currencyDeals

import (
	"deals/internal/app/domain/currencyDeals/sending/amqp"
	"deals/internal/app/domain/currencyDeals/types"
	currencyDealsLib "libs/contracts/deals/currency"
	"libs/contracts/payments"
)

func CreateCurrencyBuyDealHandler(request currencyDealsLib.CurrencyDealBuyRequest) currencyDealsLib.CurrencyDealBuyResponse {
	CreateCurrencyDealRepository(CurrencyDeal{
		UserID:   request.UserID,
		Ticker:   request.Ticker,
		Amount:   request.Amount,
		Currency: request.Currency,
		Trigger:  types.CurrencyDealTrigger(request.Trigger),
		Type:     types.BUY,
	})
	return currencyDealsLib.CurrencyDealBuyResponse{Message: "OK"}
}

func CreateCurrencySellDealHandler(request currencyDealsLib.CurrencyDealSellRequest) currencyDealsLib.CurrencyDealSellResponse {
	CreateCurrencyDealRepository(CurrencyDeal{
		UserID:   request.UserID,
		Ticker:   request.Ticker,
		Amount:   request.Amount,
		Currency: request.Currency,
		Trigger:  types.CurrencyDealTrigger(request.Trigger),
		Type:     types.SELL,
	})
	return currencyDealsLib.CurrencyDealSellResponse{Message: "OK"}
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
	deals := GetAllCurrencyDealsByTickerRepository(request.Ticker)
	for _, deal := range deals {
		if (deal.Trigger == types.UP && request.Currency > deal.Currency) ||
			(deal.Trigger == types.DOWN && request.Currency < deal.Currency) {
			if deal.Type == types.BUY {
				CurrencyDealBuyHandler(deal, request)
			} else {
				CurrencyDealSellHandler(deal, request)
			}
		}
	}
}

func CurrencyDealBuyHandler(deal CurrencyDeal, tickerChange currencyDealsLib.IncomingCurrencyChangeRequest) {
	userRUBBalance := amqp.GetUserBalances(payments.GetBalancesRequest{UserID: deal.UserID, Ticker: "RUB"})

	cost := deal.Amount * tickerChange.Currency
	if userRUBBalance.Amount < cost {
		ChangeCurrencyDealStatusRepository(deal.ID, true, "Недостаточно средств для проведения сделки.")
		return
	}

	updatedRUBBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: deal.UserID,
		Ticker: "RUB",
		Amount: -cost,
	})
	if updatedRUBBalance.Status == false {
		ChangeCurrencyDealStatusRepository(deal.ID, false,
			"Невозможно обновить рублевой баланс пользователя.")
		return
	}

	updatedTickerBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: deal.UserID,
		Ticker: deal.Ticker,
		Amount: deal.Amount,
	})

	if updatedTickerBalance.Status == false {
		amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
			UserID: deal.UserID,
			Ticker: "RUB",
			Amount: cost,
		})
		ChangeCurrencyDealStatusRepository(deal.ID, true,
			"Невозможно осуществить ролл-бек баланса пользователя. Ожидается обращение в тех. поддержку.")
		return
	}

	ChangeCurrencyDealStatusRepository(deal.ID, true, "Сделка успешно проведена.")
}

func CurrencyDealSellHandler(deal CurrencyDeal, tickerChange currencyDealsLib.IncomingCurrencyChangeRequest) {
	userBalance := amqp.GetUserBalances(payments.GetBalancesRequest{UserID: deal.UserID, Ticker: deal.Ticker})

	if userBalance.Amount < deal.Amount {
		ChangeCurrencyDealStatusRepository(deal.ID, true, "Недостаточно средств для проведения сделки.")
		return
	}

	cost := deal.Amount * tickerChange.Currency

	updatedTickerBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: deal.UserID,
		Ticker: deal.Ticker,
		Amount: -deal.Amount,
	})

	if updatedTickerBalance.Status == false {
		ChangeCurrencyDealStatusRepository(deal.ID, false,
			"Невозможно обновить рублевой баланс пользователя.")
		return
	}

	updatedRUBBalance := amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
		UserID: deal.UserID,
		Ticker: "RUB",
		Amount: cost,
	})

	if updatedRUBBalance.Status == false {
		amqp.UpdateUserBalances(payments.UpdateBalanceRequest{
			UserID: deal.UserID,
			Ticker: deal.Ticker,
			Amount: deal.Amount,
		})
		ChangeCurrencyDealStatusRepository(deal.ID, true,
			"Невозможно осуществить ролл-бек баланса пользователя. Ожидается обращение в тех. поддержку.")
		return
	}

	ChangeCurrencyDealStatusRepository(deal.ID, true, "Сделка успешно проведена.")
}
