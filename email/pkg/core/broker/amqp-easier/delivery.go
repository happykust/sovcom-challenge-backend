package amqp_easier

import (
	"email/pkg/core/broker/amqp-easier/constants"
	logger "email/pkg/logging"
	LoggerTypes "email/pkg/logging/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumerConstructor(connName string, exchangeName string, exchangeType string, routingKey string, queueName string) (<-chan amqp.Delivery, *amqp.Channel, *amqp.Connection) {
	amqpChannel, conn := ConnectToRabbit(connName)

	err := amqpChannel.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, constants.ERROR_FAILED_TO_DECLARE_EXCHANGE, err)
	}

	queue, err := amqpChannel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, constants.ERROR_FAILED_TO_DECLARE_QUEUE, err)
	}

	amqpChannel.QueueBind(queue.Name, routingKey, exchangeName, false, nil)

	messageChannel, err := amqpChannel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, constants.ERROR_FAILED_TO_REGISTER_CONSUMER, err)
	}

	return messageChannel, amqpChannel, conn
}
