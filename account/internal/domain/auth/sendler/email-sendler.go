package sendler

import (
	amqp_easier "account/pkg/core/broker/amqp-easier"
	"libs/contracts/email"
)

func SendEmail(jsonObj []byte) {
	connName := "test"
	exchangeName := email.EmailExchange
	exchangeType := "topic"
	routingKey := email.Topic
	body := jsonObj

	amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)

}
