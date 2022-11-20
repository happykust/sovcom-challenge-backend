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

func PromocodesUpdateConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(promocodes.UpdateConsumerName,
		promocodes.PromocodesExchange, "topic", promocodes.UpdateTopic, promocodes.UpdateQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "[Promocodes | Update consumer] Waiting for messages...", nil)
		for d := range messageChannel {

			promocodeRequest := &promocodes.UpdateRequest{}
			err := json.Unmarshal(d.Body, promocodeRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Promocodes | Update consumer] Can't unmarshal incoming body", err)
			}

			UpdatePromocode := promocodesInternal.PromocodesUpdateHandler(promocodesInternal.Promocode{
				Promocode:            promocodeRequest.Promocode,
				Ticker:               promocodeRequest.Ticker,
				Amount:               promocodeRequest.Amount,
				ActivationCountLimit: promocodeRequest.ActivationCountLimit,
			})
			response, err := json.Marshal(&promocodes.UpdateResponse{Promocode: UpdatePromocode.Promocode,
				Ticker: UpdatePromocode.Ticker, Amount: UpdatePromocode.Amount,
				ActivationCountLimit: UpdatePromocode.ActivationCountLimit})

			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Promocodes | Update consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
