package currencyDeals

import "deals/pkg/core/database"

func GetAllCurrencyDealsRepository() []CurrencyDeal {
	var deals []CurrencyDeal
	database.PG.Find(&deals)
	return deals
}

func GetAllCurrencyDealsByTickerRepository(ticker string) []CurrencyDeal {
	var deals []CurrencyDeal
	database.PG.Where("ticker = ?", ticker).Find(&deals)
	return deals
}

func GetCurrencyDealRepository(UserID uint) []CurrencyDeal {
	var deals []CurrencyDeal
	database.PG.Where("user_id = ?", UserID).Find(&deals)
	return deals
}

func CreateCurrencyDealRepository(deal CurrencyDeal) CurrencyDeal {
	database.PG.Create(&deal)
	return deal
}

func DeleteCurrencyDealRepository(dealID uint) uint {
	database.PG.Where("id = ?", dealID).Delete(&CurrencyDeal{})
	return dealID
}

func ChangeCurrencyDealStatusRepository(dealID uint, status bool, message string) uint {
	database.PG.Model(&CurrencyDeal{}).Where("id = ?", dealID).
		Update("status", status).
		Update("message", message).
		Update("tried", true)
	return dealID
}
