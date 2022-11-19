package amqp_easier

import (
	"currency-parser/pkg/core/broker/amqp-easier/constants"
	"currency-parser/pkg/core/config"
	logger "currency-parser/pkg/logging"
	LoggerTypes "currency-parser/pkg/logging/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbit(connName string) (*amqp.Channel, *amqp.Connection) {
	configConnect := amqp.Config{Properties: amqp.NewConnectionProperties()}
	configConnect.Properties.SetClientConnectionName(connName)
	conn, err := amqp.DialConfig(config.GetAMQPUri(), configConnect)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, constants.ERROR_FAILED_TO_CONNECT_TO_AMQP, err)
	}

	amqpChannel, err := conn.Channel()
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, constants.ERROR_FAILED_TO_OPEN_CHANNEL, err)
	}

	return amqpChannel, conn
}
