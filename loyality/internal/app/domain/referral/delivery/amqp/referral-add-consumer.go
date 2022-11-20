package amqp

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"libs/contracts/loyality/referral"
	referralInternal "loyality/internal/app/domain/referral"
	amqp_easier "loyality/pkg/core/broker/amqp-easier"
	logger "loyality/pkg/logging"
	LoggerTypes "loyality/pkg/logging/types"
)

func ReferralAddConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(referral.AddConsumerName,
		referral.ReferralExchange, "topic", referral.AddTopic, referral.AddQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "[Referral | Add consumer] Waiting for messages...", nil)
		for d := range messageChannel {

			ReferralRequest := &referral.AddRequest{}
			err := json.Unmarshal(d.Body, ReferralRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Referral | Add consumer] Can't unmarshal incoming body", err)
			}
			addedReferral := referralInternal.ReferralIncreaseReferralCountHandler(*ReferralRequest)
			response, err := json.Marshal(&referral.AddResponse{
				Message: addedReferral,
			})
			
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Referral | Add consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
