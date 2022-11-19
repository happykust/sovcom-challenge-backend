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

func CurrencyDealSellConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(CurrencyDeals.SellConsumerName,
		CurrencyDeals.CurrencyDealsExchange, "topic", CurrencyDeals.SellTopic, CurrencyDeals.SellQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Deals | CurrencyDealSell consumer] Waiting for messages...")
		for d := range messageChannel {

			dealRequest := &CurrencyDeals.CurrencyDealSellRequest{}
			err := json.Unmarshal(d.Body, dealRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Deals | CurrencyDealSell consumer] Can't unmarshal incoming body", err)
			}

			createdDeal := currencyDeals.CreateCurrencySellDealHandler(*dealRequest)
			response, err := json.Marshal(&createdDeal)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType: "text/plain",
					Body:        response,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Deals | CurrencyDealSell consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
