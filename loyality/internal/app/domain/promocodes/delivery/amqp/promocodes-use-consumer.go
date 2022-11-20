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

func PromocodesUseConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(promocodes.UseConsumerName,
		promocodes.PromocodesExchange, "topic", promocodes.UseTopic, promocodes.UseQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "[Promocodes | Use consumer] Waiting for messages...", nil)
		for d := range messageChannel {

			promocodeRequest := &promocodes.UseRequest{}
			err := json.Unmarshal(d.Body, promocodeRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Promocodes | Use consumer] Can't unmarshal incoming body", err)
			}

			UsePromocode := promocodesInternal.PromocodesUseHandler(*promocodeRequest)
			response, err := json.Marshal(&promocodes.UseResponse{Message: UsePromocode})

			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType: "text/plain",
					Body:        response,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Promocodes | Use consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
