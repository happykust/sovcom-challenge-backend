package amqp

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"libs/contracts/payments"
	"payments/internal/app/domain/transactions"
	amqp_easier "payments/pkg/core/broker/amqp-easier"
	logger "payments/pkg/logging"
	LoggerTypes "payments/pkg/logging/types"
)

func CreateConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(payments.CreateTransactionConsumerName,
		payments.PaymentsExchange, "topic", payments.CreateTransactionTopic, payments.CreateTransactionQueueName)

	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "Waiting for messages...", nil)
		for d := range messageChannel {

			createTransaction := payments.CreateTransactionRequest{}
			err := json.Unmarshal(d.Body, &createTransaction)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "Error while unmarshalling create transaction request", err)
				continue
			}
			createdTransaction := transactions.CreateTransaction(createTransaction)
			t, err := json.Marshal(createdTransaction)

			fmt.Println(string(d.Body))
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "Error while publishing created transaction response", err)
				}
			}
		}
	}()
	<-stopChan
}

func GetConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(payments.GetTransactionConsumerName,
		payments.PaymentsExchange, "topic", payments.GetTransactionTopic, payments.GetTransactionQueueName)
	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "Waiting for messages...", nil)
		for d := range messageChannel {
			getTransaction := payments.PaymentsTransactionsGet{}
			err := json.Unmarshal(d.Body, &getTransaction)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "Error while unmarshalling get transaction request", err)
				continue
			}
			transaction := transactions.GetTransaction(getTransaction.TransactionUUID)
			t, err := json.Marshal(transaction)
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "Error while publishing get transaction response", err)
				}
			}
		}
	}()
	<-stopChan
}

func UpdateTransactionStatusConsumer() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(payments.UpdateTransactionStatusConsumerName,
		payments.PaymentsExchange, "topic", payments.UpdateTransactionStatusTopic, payments.UpdateTransactionStatusQueueName)
	defer amqpChannel.Close()
	defer conn.Close()
	stopChan := make(chan bool)
	go func() {
		logger.Log(LoggerTypes.INFO, "Waiting for messages...", nil)
		for d := range messageChannel {
			updateTransactionStatus := payments.PaymentsTransactionsUpdateRequest{}
			err := json.Unmarshal(d.Body, &updateTransactionStatus)
			if err != nil {
				logger.Log(LoggerTypes.CRITICAL, "Error while unmarshalling update transaction status request", err)
				continue
			}
			transactions.UpdateTransactionStatus(updateTransactionStatus.TransactionUUID,
				payments.TransactionStatus(updateTransactionStatus.TransactionStatus))
			t, err := json.Marshal(payments.PaymentsTransactionsUpdateResponse{Message: "Transaction status updated"})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					logger.Log(LoggerTypes.CRITICAL, "Error while publishing update transaction status response", err)
				}
			}
		}
	}()
	<-stopChan
}
