package amqp

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"libs/contracts/payments"
	"payments/internal/app/domain/balance"
	amqp_easier "payments/pkg/core/broker/amqp-easier"
	logger "payments/pkg/logging"
	LoggerTypes "payments/pkg/logging/types"
)

func UserCreatedEventCreateBalance() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(payments.CreateBalanceQueueName, payments.PaymentsExchange, "topic", payments.CreateBalanceRoutingKey, payments.CreateBalanceQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			createBalance := payments.CreateBalanceUserRequest{}
			err := json.Unmarshal(d.Body, &createBalance)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "Error while unmarshalling create balance request", err)
			}
			balance := balance.CreateUserEvent(createBalance.UserID)
			t, err := json.Marshal(&payments.CreateBalanceUserResponse{BalanceId: balance.ID})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "Error while publishing create balance response", err)
				}
			}
		}

	}()
	<-stopChan
}

// create user balance

// create user rub waller

//func TestConsumer() {
//
//	messageChannel, amqpChannel, conn := amqp_easier("mem", "priver", "topic", "gdfdfg", "testow")
//
//	defer amqpChannel.Close()
//	defer conn.Close()
//	stopChan := make(chan bool)
//	go func() {
//		fmt.Println("waiting for message")
//		for d := range messageChannel {
//			fmt.Println("waiting for message")
//			// createUser := &account.RegisterRequest{}
//			// err := json.Unmarshal(d.Body, createUser)
//			// if err != nil {
//			//	 logger.Log(LoggerTypes.CRITICAL, constants.ERROR_FAILED_TO_REGISTER_CONSUMER, err)
//			// }
//			// register := balance.SingUp(createUser.Email, createUser.Password, createUser.FullName, createUser.StudentId)
//			// t, err := json.Marshal(&account.RegisterResponse{Message: register.Message, AccessToken: register.AccessToken, RefreshToken: register.RefreshToken})
//			//
//			fmt.Println(string(d.Body))
//			if len(d.ReplyTo) != 0 {
//				ctx := context.Background()
//				msg := amqp.Publishing{
//					ContentType:   "text/plain",
//					Body:          d.Body,
//					CorrelationId: d.CorrelationId,
//				}
//
//				err := amqpChannel.PublishWithContext(ctx, "", "", false, false, msg)
//
//				if err != nil {
//					logger.Log(LoggerTypes.CRITICAL, "Error publishing message", err)
//				}
//			}
//
//		}
//
//	}()
//	<-stopChan
//}
