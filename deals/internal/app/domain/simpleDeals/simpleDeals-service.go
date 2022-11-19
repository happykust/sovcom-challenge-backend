package simpleDeals

import (
	types2 "deals/internal/app/domain/simpleDeals/types"
	"deals/pkg/core/config"
	"deals/pkg/core/database"
	logger "deals/pkg/logging"
	LoggerTypes "deals/pkg/logging/types"
	"encoding/json"
	"libs/contracts/currency"
)

func GetTickerCurrency(request currency.ReadTickerRequest) currency.ReadTickerResponse {
	var response currency.ReadTickerResponse

	lastKline := database.Redis.Get(config.RedisLastCurrenciesTag + ":" + request.TickerGroup).Val()
	if lastKline == "" {
		return response
	}

	var lastKlineStruct types2.GetCurrencyJSONResponse
	err := json.Unmarshal([]byte(lastKline), &lastKlineStruct)
	if err != nil {
		logger.Log(LoggerTypes.ERROR, "[Deals | SimpleDeals] Error unmarshal last kline", err)
		return response
	}

	response.Currency = lastKlineStruct.Kline.ClosePrice
	return response
}
