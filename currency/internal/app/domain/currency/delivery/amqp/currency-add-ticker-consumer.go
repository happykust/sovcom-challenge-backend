package amqp

import (
	currencyInternal "currency/internal/app/domain/currency"
	amqp_easier "currency/pkg/core/broker/amqp-easier"
	logger "currency/pkg/logging"
	LoggerTypes "currency/pkg/logging/types"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"libs/contracts/currency"
)

func CurrencyAddTickerConsumer() {

	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(currency.AddTickerConsumerName,
		currency.CurrencyExchange, "topic", currency.AddTickerTopic, currency.AddTickerQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Currency | Add ticker consumer] Waiting for messages...")
		for d := range messageChannel {

			tickerRequest := &currency.AddTickerRequest{}
			err := json.Unmarshal(d.Body, tickerRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Currency | Add ticker consumer] Can't unmarshal incoming body", err)
			}
			createdTicker := currencyInternal.CurrencyAddTickerHandler(*tickerRequest)
			response, err := json.Marshal(&currency.AddTickerResponse{
				Ticker: createdTicker.Ticker,
				Group:  createdTicker.Group,
			})

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Currency | Add ticker consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
