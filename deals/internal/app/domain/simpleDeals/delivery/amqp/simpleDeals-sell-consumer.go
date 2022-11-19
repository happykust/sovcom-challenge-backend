package amqp

import (
	simpleDealsInternal "deals/internal/app/domain/simpleDeals"
	amqp_easier "deals/pkg/core/broker/amqp-easier"
	logger "deals/pkg/logging"
	LoggerTypes "deals/pkg/logging/types"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	simpleDeals "libs/contracts/deals/simple"
)

func SimpleDealSellConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(simpleDeals.SellConsumerName,
		simpleDeals.SimpleDealsExchange, "topic", simpleDeals.SellTopic, simpleDeals.SellQueueName)

	defer amqpChannel.Close()
	defer conn.Close()

	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Deals | SimpleDealSell consumer] Waiting for messages...")
		for d := range messageChannel {

			dealRequest := &simpleDeals.SimpleDealSellRequest{}
			err := json.Unmarshal(d.Body, dealRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL,
					"[Deals | SimpleDealSell consumer] Can't unmarshal incoming body", err)
			}

			createdDeal := simpleDealsInternal.SimpleDealSellHandler(*dealRequest)
			response, err := json.Marshal(&createdDeal)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType: "text/plain",
					Body:        response,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL,
						"[Deals | SimpleDealSell consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
