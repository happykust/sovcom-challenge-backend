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

func CurrencyReadConsumer() {

	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(currency.ReadConsumerName,
		currency.CurrencyExchange, "topic", currency.ReadTopic, currency.ReadQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Currency | Read consumer] Waiting for messages...")
		for d := range messageChannel {

			tickerRequest := &currency.ReadRequest{}
			err := json.Unmarshal(d.Body, tickerRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Currency | Read consumer] Can't unmarshal incoming body", err)
			}
			createdTicker := currencyInternal.CurrencyReadHandler()
			response, err := json.Marshal(createdTicker)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Currency | Read consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
