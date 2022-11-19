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

func CurrencyReadGroupConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(currency.ReadGroupConsumerName,
		currency.CurrencyExchange, "topic", currency.ReadGroupTopic, currency.ReadGroupQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Currency | ReadGroup consumer] Waiting for messages...")
		for d := range messageChannel {

			GroupRequest := &currency.ReadGroupsRequest{}
			err := json.Unmarshal(d.Body, GroupRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Currency | ReadGroup consumer] Can't unmarshal incoming body", err)
			}
			ReadGroupdTicker := currencyInternal.CurrencyReadTypeHandler()
			response, err := json.Marshal(ReadGroupdTicker)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Currency | ReadGroup consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
