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

func PromocodesDeleteConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(promocodes.DeleteConsumerName,
		promocodes.PromocodesExchange, "topic", promocodes.DeleteTopic, promocodes.DeleteQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "[Promocodes | Delete consumer] Waiting for messages...", nil)
		for d := range messageChannel {

			promocodeRequest := &promocodes.DeleteRequest{}
			err := json.Unmarshal(d.Body, promocodeRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Promocodes | Delete consumer] Can't unmarshal incoming body", err)
			}

			deletedPromocodeName := promocodesInternal.PromocodesDeleteHandler(*promocodeRequest)
			response, err := json.Marshal(&promocodes.DeleteResponse{Promocode: deletedPromocodeName})
			
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType: "text/plain",
					Body:        response,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Promocodes | Delete consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
