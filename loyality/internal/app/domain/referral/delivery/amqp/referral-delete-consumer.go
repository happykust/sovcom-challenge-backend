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

func ReferralDeleteConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(referral.DeleteConsumerName,
		referral.ReferralExchange, "topic", referral.DeleteTopic, referral.DeleteQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "[Referral | Delete consumer] Waiting for messages...", nil)
		for d := range messageChannel {

			ReferralRequest := &referral.DeleteRequest{}
			err := json.Unmarshal(d.Body, ReferralRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Referral | Delete consumer] Can't unmarshal incoming body", err)
			}
			deletedReferral := referralInternal.ReferralDeleteHandler(*ReferralRequest)
			response, err := json.Marshal(&referral.DeleteResponse{
				UUID: deletedReferral,
			})

			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Referral | Delete consumer] Error publishing message", err)
				}
			}

		}

	}()
	<-stopChan
}
