package simpleDeals

import (
	"deals/internal/app/domain/simpleDeals/types"
	"deals/pkg/core/config"
	"deals/pkg/core/database"
	logger "deals/pkg/logging"
	LoggerTypes "deals/pkg/logging/types"
	"encoding/json"
	"fmt"
	"libs/contracts/currency"
	"strconv"
)

func GetTickerCurrency(request currency.ReadTickerRequest) currency.ReadTickerResponse {
	var response currency.ReadTickerResponse

	lastKline := database.Redis.Get(config.RedisLastCurrenciesTag + ":" + request.TickerGroup).Val()
	fmt.Println(request.TickerGroup, lastKline)
	if lastKline == "" {
		return response
	}

	var lastKlineStruct types.GetCurrencyJSONResponse
	err := json.Unmarshal([]byte(lastKline), &lastKlineStruct)
	if err != nil {
		logger.Log(LoggerTypes.ERROR, "[Deals | SimpleDeals] Error unmarshal last kline", err)
		return response
	}

	response.Currency, _ = strconv.ParseFloat(lastKlineStruct.Kline.ClosePrice, 64)
	response.TickerFrom = lastKlineStruct.TickerFrom
	response.TickerTo = lastKlineStruct.TickerTo
	return response
}
