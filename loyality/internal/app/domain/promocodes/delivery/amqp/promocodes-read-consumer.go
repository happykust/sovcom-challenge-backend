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

func PromocodesReadConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(promocodes.ReadConsumerName,
		promocodes.PromocodesExchange, "topic", promocodes.ReadTopic, promocodes.ReadQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "[Promocodes | Read consumer] Waiting for messages...", nil)
		for d := range messageChannel {

			promocodeRequest := &promocodes.ReadRequest{}
			err := json.Unmarshal(d.Body, promocodeRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Promocodes | Read consumer] Can't unmarshal incoming body", err)
			}

			ReadPromocode := promocodesInternal.PromocodesReadHandler(*promocodeRequest)
			response, err := json.Marshal(&promocodes.ReadResponse{Promocode: ReadPromocode.Promocode,
				Ticker: ReadPromocode.Ticker, Amount: ReadPromocode.Amount,
				ActivationCountLimit: ReadPromocode.ActivationCountLimit})
			
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType: "text/plain",
					Body:        response,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Promocodes | Read consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
