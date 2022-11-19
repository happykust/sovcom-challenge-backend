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

func CurrencyCreateGroupConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(currency.CreateGroupConsumerName,
		currency.CurrencyExchange, "topic", currency.CreateGroupTopic, currency.CreateGroupQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Currency | CreateGroup consumer] Waiting for messages...")
		for d := range messageChannel {

			GroupRequest := &currency.CreateGroupRequest{}
			err := json.Unmarshal(d.Body, GroupRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Currency | CreateGroup consumer] Can't unmarshal incoming body", err)
			}
			CreateGroupTicker := currencyInternal.CurrencyCreateTypeHandler(currencyInternal.Group{
				Title: GroupRequest.Title})
			response, err := json.Marshal(CreateGroupTicker)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Currency | CreateGroup consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
