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

func SimpleDealBuyConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(simpleDeals.BuyConsumerName,
		simpleDeals.SimpleDealsExchange, "topic", simpleDeals.BuyTopic, simpleDeals.BuyQueueName)

	defer amqpChannel.Close()
	defer conn.Close()

	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Deals | SimpleDealBuy consumer] Waiting for messages...")
		for d := range messageChannel {

			dealRequest := &simpleDeals.SimpleDealBuyRequest{}
			err := json.Unmarshal(d.Body, dealRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL,
					"[Deals | SimpleDealBuy consumer] Can't unmarshal incoming body", err)
			}

			createdDeal := simpleDealsInternal.SimpleDealBuyHandler(*dealRequest)
			response, err := json.Marshal(&createdDeal)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL,
						"[Deals | SimpleDealBuy consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
