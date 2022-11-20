package amqp

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"libs/contracts/loyality/promocodes"
	promocodesInternal "loyality/internal/app/domain/promocodes"
	amqp_easier "loyality/pkg/core/broker/amqp-easier"
	logger "loyality/pkg/logging"
	LoggerTypes "loyality/pkg/logging/types"
)

func PromocodesCreateConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(promocodes.CreateConsumerName,
		promocodes.PromocodesExchange, "topic", promocodes.CreateTopic, promocodes.CreateQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "[Promocodes | Create consumer] Waiting for messages...", nil)
		for d := range messageChannel {

			promocodeRequest := &promocodes.CreateRequest{}
			err := json.Unmarshal(d.Body, promocodeRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Promocodes | Create consumer] Can't unmarshal incoming body", err)
			}

			createdPromocode := promocodesInternal.PromocodesCreateHandler(*promocodeRequest)
			response, err := json.Marshal(&promocodes.CreateResponse{Promocode: createdPromocode.Promocode,
				Ticker: createdPromocode.Ticker, Amount: createdPromocode.Amount,
				ActivationCountLimit: createdPromocode.ActivationCountLimit})

			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Promocodes | Create consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
