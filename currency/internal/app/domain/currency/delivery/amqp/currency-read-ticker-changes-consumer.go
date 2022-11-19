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

func CurrencyReadTickerChangesConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(currency.ReadTickerChangesConsumerName,
		currency.CurrencyExchange, "topic", currency.ReadTickerChangesTopic, currency.ReadTickerChangesQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Currency | ReadTickerChanges consumer] Waiting for messages...")
		for d := range messageChannel {

			TickerChangesRequest := &currency.ReadTickerChangesRequest{}
			err := json.Unmarshal(d.Body, TickerChangesRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Currency | ReadTickerChanges consumer] Can't unmarshal incoming body", err)
			}
			ReadTickerChangesTicker := currencyInternal.CurrencyReadTickerChangesHandler(TickerChangesRequest.Ticker)
			response, err := json.Marshal(ReadTickerChangesTicker)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Currency | ReadTickerChanges consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
