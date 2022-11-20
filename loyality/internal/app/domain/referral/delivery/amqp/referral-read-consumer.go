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

func ReferralReadConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(referral.ReadConsumerName,
		referral.ReferralExchange, "topic", referral.ReadTopic, referral.ReadQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "[Referral | Read consumer] Waiting for messages...", nil)
		for d := range messageChannel {

			referralRequest := &referral.ReadRequest{}
			err := json.Unmarshal(d.Body, referralRequest)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "[Referral | Read consumer] Can't unmarshal incoming request", err)
			}
			readReferral := referralInternal.ReferralReadHandler(*referralRequest)
			depositBonus := referralInternal.ReferralMathDepositBonusHandler(readReferral.ReferralCount)
			response, err := json.Marshal(&referral.ReadResponse{UserID: readReferral.UserID, UUID: readReferral.UUID,
				Ticker: readReferral.Ticker, Amount: readReferral.Amount, ReferralCount: readReferral.ReferralCount,
				DepositBonus: depositBonus})

			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          response,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Referral | Read consumer] Error publishing message", err)
				}
			}
		}
	}()
	<-stopChan
}
