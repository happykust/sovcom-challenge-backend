package delivery

import (
	currency_ws "currency/internal/app/api/domain/currency-ws"
	amqp_easier "currency/pkg/core/broker/amqp-easier"
	logger "currency/pkg/logging"
	LoggerTypes "currency/pkg/logging/types"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"libs/contracts/currency"
)

func CurrencyUpdateConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(currency.CurrencyUpdateConsumerName,
		currency.CurrencyExchange, "topic", currency.CurrencyUpdateTopic, currency.CurrencyUpdateQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)

	go func() {
		logger.Log(LoggerTypes.INFO, "[Currency | Update consumer] Waiting for message", nil)
		for d := range messageChannel {
			currencyUpdate := &currency.CurrencyUpdateRequest{}
			err := json.Unmarshal(d.Body, currencyUpdate)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL,
					"[Currency | Update consumer] Error with unmarshaling incoming request", err)
			}

			go currency_ws.ReceiveCurrencyUpdate(*currencyUpdate)

			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          []byte(""),
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Currency | Delete consumer] Error publishing message", err)
				}
			}
		}
	}()
	<-stopChan
}
