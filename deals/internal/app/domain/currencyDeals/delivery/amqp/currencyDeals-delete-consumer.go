package amqp

import (
	"deals/internal/app/domain/currencyDeals"
	amqp_easier "deals/pkg/core/broker/amqp-easier"
	logger "deals/pkg/logging"
	LoggerTypes "deals/pkg/logging/types"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	CurrencyDeals "libs/contracts/deals/currency"
)

func CurrencyDealDeleteConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(CurrencyDeals.DeleteConsumerName,
		CurrencyDeals.CurrencyDealsExchange, "topic", CurrencyDeals.DeleteTopic, CurrencyDeals.DeleteQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Deals | CurrencyDealDelete consumer] Waiting for messages...")
		for d := range messageChannel {

			dealRequest := &CurrencyDeals.CurrencyDealDeleteRequest{}
			err := json.Unmarshal(d.Body, dealRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Deals | CurrencyDealDelete consumer] Can't unmarshal incoming body", err)
			}

			createdDeal := currencyDeals.DeleteCurrencyDealHandler(*dealRequest)
			response, err := json.Marshal(&createdDeal)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Deals | CurrencyDealDelete consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
