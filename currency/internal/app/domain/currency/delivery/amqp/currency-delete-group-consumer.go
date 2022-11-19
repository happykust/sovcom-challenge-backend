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

func CurrencyDeleteGroupConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(currency.DeleteGroupConsumerName,
		currency.CurrencyExchange, "topic", currency.DeleteGroupTopic, currency.DeleteGroupQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Currency | DeleteGroup consumer] Waiting for messages...")
		for d := range messageChannel {

			GroupRequest := &currency.DeleteGroupRequest{}
			err := json.Unmarshal(d.Body, GroupRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Currency | DeleteGroup consumer] Can't unmarshal incoming body", err)
			}
			DeleteGroupdTicker := currencyInternal.CurrencyDeleteTypeHandler(GroupRequest.Title)
			response, err := json.Marshal(DeleteGroupdTicker)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Currency | DeleteGroup consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
