package amqp

import (
	"email/internal/app/domain/smtp"
	amqp_easier "email/pkg/core/broker/amqp-easier"
	logger "email/pkg/logging"
	LoggerTypes "email/pkg/logging/types"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"libs/contracts/email"
)

func EmailConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(email.ConsumerName, email.EmailExchange,
		"topic", email.Topic, email.QueueName)
	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)

	go func() {
		fmt.Println("waiting for message")
		for d := range messageChannel {
			sendEmail := &email.Request{}
			err := json.Unmarshal(d.Body, sendEmail)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "ERR", err)
			}
			sendEmailAction := smtp.SendEmail(sendEmail.Email, sendEmail.Subject, sendEmail.Body)
			response, err := json.Marshal(&email.Response{Message: sendEmailAction})

			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType: "text/plain",
					Body:        response,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Email | Email consumer] Error publishing message", err)
				}
			}
		}
	}()
	<-stopChan
}

func EmailByUserIDConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(email.SendEmailByUserIDConsumerName, email.EmailExchange,
		"topic", email.SendEmailByUserIDTopic, email.SendEmailByUserIDQueueName)
	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)

	go func() {
		fmt.Println("waiting for message")
		for d := range messageChannel {
			sendEmail := &email.SendEmailByUserIDRequest{}
			err := json.Unmarshal(d.Body, sendEmail)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "ERR", err)
			}

			sendEmailAction := smtp.SendEmailByUserID(sendEmail.Subject, sendEmail.Body, sendEmail.UserID)
			response, err := json.Marshal(&email.Response{Message: sendEmailAction})

			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType: "text/plain",
					Body:        response,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "[Email | Email consumer] Error publishing message", err)
				}
			}
		}
	}()
	<-stopChan
}
