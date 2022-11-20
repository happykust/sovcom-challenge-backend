package amqp

import (
	"deals/internal/app/domain/currencyDeals"
	amqp_easier "deals/pkg/core/broker/amqp-easier"
	logger "deals/pkg/logging"
	LoggerTypes "deals/pkg/logging/types"
	"encoding/json"
	"fmt"
	CurrencyDeals "libs/contracts/deals/currency"
)

func CurrencyDealCurrencyChangeConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(CurrencyDeals.CurrencyChangeConsumerName,
		CurrencyDeals.CurrencyDealsExchange, "topic", CurrencyDeals.CurrencyChangeTopic,
		CurrencyDeals.CurrencyChangeQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		fmt.Println("[Deals | CurrencyDealCurrencyChange consumer] Waiting for messages...")
		for d := range messageChannel {

			dealRequest := &CurrencyDeals.IncomingCurrencyChangeRequest{}
			err := json.Unmarshal(d.Body, dealRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL,
					"[Deals | CurrencyDealCurrencyChange consumer] Can't unmarshal incoming body", err)
			}

			go currencyDeals.ExecuteCurrencyDealsHandler(*dealRequest)

			fmt.Println(string(d.Body))
		}

	}()
	<-stopChan
}
