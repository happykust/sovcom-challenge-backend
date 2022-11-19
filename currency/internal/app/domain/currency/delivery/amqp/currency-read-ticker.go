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

func CurrencyReadTickerConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(currency.ReadTickerConsumerName,
		currency.CurrencyExchange, "topic", currency.ReadTickerTopic, currency.ReadTickerQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Currency | ReadTicker consumer] Waiting for messages...")
		for d := range messageChannel {

			tickerRequest := &currency.ReadTickerRequest{}
			err := json.Unmarshal(d.Body, tickerRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Currency | ReadTicker consumer] Can't unmarshal incoming body", err)
			}
			ReadTickerTicker := currencyInternal.CurrencyReadTickerHandler(*tickerRequest)
			response, err := json.Marshal(ReadTickerTicker)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Currency | ReadTicker consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
