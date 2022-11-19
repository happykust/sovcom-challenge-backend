package currency

import (
	"context"
	"currency-parser/internal/app/sending/amqp"
	"currency-parser/pkg/core/config"
	"currency-parser/pkg/core/database"
	logger "currency-parser/pkg/logging"
	LoggerTypes "currency-parser/pkg/logging/types"
	"encoding/json"
	"io"
	"libs/contracts/currency"
	"net/http"
	"strconv"
	"time"
)

func Consumer(tickerFrom string, ticketTo string) {
	tickerFullChanged := false
	tickerFull := tickerFrom + ticketTo
	logger.Log(LoggerTypes.INFO, "[Currency-parser | Currency | "+tickerFull+"] Start parsing.", nil)

	var unixTimeStart int64
	unixTimeEnd := time.Now().Unix()

	allRows := database.Redis.HGetAll(context.Background(), TickersGroupName+":"+tickerFull).Val()
	allKeys := database.Redis.HKeys(context.Background(), TickersGroupName+":"+tickerFull).Val()
	var uniTimeMax = int64(0)

	for _, key := range allKeys {
		value := allRows[key]
		parsedValue := TickerDynamicJSONOutcoming{}
		err := json.Unmarshal([]byte(value), &parsedValue)
		if err != nil {
			logger.Log(LoggerTypes.ERROR, "[Currency-parser | Currency | "+tickerFull+"] Error while unmarshaling redis stored message.", err)
		}
		unixTime := parsedValue.Timestamp
		if err != nil {
			logger.Log(LoggerTypes.ERROR, "[Currency-parser | Currency | "+tickerFull+"] Error while parsing unix time.", err)
		}
		if unixTime > uniTimeMax {
			uniTimeMax = unixTime
		}
	}

	unixTimeStart = uniTimeMax

	logger.Log(LoggerTypes.INFO, "[Currency-parser | Currency | "+tickerFull+"] Last stored unix time: "+strconv.FormatInt(unixTimeStart, 10), nil)

	if len(allKeys) == 0 {
		logger.Log(LoggerTypes.INFO, "[Currency-parser | Currency | "+tickerFull+"] No data in redis. Start parsing from ~2 years ago.", nil)
		unixTimeStart = time.Now().Unix() - 86400*364*2
	} else {
		logger.Log(LoggerTypes.INFO, "[Currency-parser | Currency | "+tickerFull+"] Start parsing from last parsed date.", nil)
	}

	for {
		unixTimeEnd = time.Now().Unix()
		connectionString := "https://query1.finance.yahoo.com/v8/finance/chart/" + tickerFull + "=X?symbol=" + tickerFull + "%3DX&period1=" + strconv.FormatInt(unixTimeStart, 10) + "&period2=" + strconv.FormatInt(unixTimeEnd, 10) + "&useYfid=true&interval=1h&includePrePost=true&events=div%7Csplit%7Cearn&lang=en-US&region=US&crumb=undefined&corsDomain=finance.yahoo.com"
		req, err := http.Get(connectionString)
		if err != nil {
			logger.Log(LoggerTypes.ERROR, "[Currency-parser | Currency | "+tickerFull+"] Error while creating request.", err)
			continue
		}

		tickerDynamic := TickerDynamicJSONIncoming{}

		byteResponse, err := io.ReadAll(req.Body)
		if err != nil {
			logger.Log(LoggerTypes.ERROR, "[Currency-parser | Currency | "+tickerFull+"] Error while reading response from server.", err)
			continue
		}

		err = json.Unmarshal(byteResponse, &tickerDynamic)
		if err != nil {
			logger.Log(LoggerTypes.ERROR, "[Currency-parser | Currency | "+tickerFull+"] Error while unmarshalling response.", err)
			continue
		}

		if len(tickerDynamic.Chart.Result) == 0 {
			if tickerFullChanged {
				logger.Log(LoggerTypes.INFO, "[Currency-parser | Currency | "+tickerFull+"] No data in response. Stop parsing.", nil)
				break
			} else {
				tickerFull = ticketTo + tickerFrom
				tickerFullChanged = true
				continue
			}
		}

		timeStampLen := len(tickerDynamic.Chart.Result[0].Timestamp)
		for i := 0; i < timeStampLen; i++ {
			tickerDynamicJSONOutcoming := TickerDynamicJSONOutcoming{
				Timestamp: tickerDynamic.Chart.Result[0].Timestamp[i],
				Open:      tickerDynamic.Chart.Result[0].Indicators.Quote[0].Open[i],
				High:      tickerDynamic.Chart.Result[0].Indicators.Quote[0].High[i],
				Low:       tickerDynamic.Chart.Result[0].Indicators.Quote[0].Low[i],
				Close:     tickerDynamic.Chart.Result[0].Indicators.Quote[0].Close[i],
				Volume:    tickerDynamic.Chart.Result[0].Indicators.Quote[0].Volume[i],
			}
			byteResponseJSON, err := json.Marshal(tickerDynamicJSONOutcoming)
			if err != nil {
				logger.Log(LoggerTypes.ERROR, "[Currency-parser | Currency | "+tickerFull+"] Error while marshalling response.", err)
				continue
			}
			database.Redis.HSet(context.Background(), TickersGroupName+":"+tickerFull,
				tickerDynamicJSONOutcoming.Timestamp, byteResponseJSON)
			amqp.SendCurrencyUpdate(currency.CurrencyUpdateRequest{TickerGroup: tickerFull, Data: tickerDynamicJSONOutcoming})
			logger.Log(LoggerTypes.INFO, "[Currency-parser | Currency | "+tickerFull+"] Parsed data for "+time.Unix(tickerDynamicJSONOutcoming.Timestamp, 0).String(), nil)
		}
		req.Body.Close()
		database.Redis.SAdd(context.Background(), config.RedisTickersGroupsSet, tickerFull)
		unixTimeStart = time.Now().Unix()
		time.Sleep(time.Hour + time.Second)
	}
}
