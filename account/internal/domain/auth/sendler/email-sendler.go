package sendler

import (
	amqp_easier "account/pkg/core/broker/amqp-easier"
	"libs/contracts/email"
	"libs/contracts/payments"
)

func SendEmail(jsonObj []byte) {
	connName := "test"
	exchangeName := email.EmailExchange
	exchangeType := "topic"
	routingKey := email.Topic
	body := jsonObj

	amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
}

func SendPayments(jsonObj []byte) []byte {
	connName := "tesst"
	exchangeName := payments.PaymentsExchange
	exchangeType := "topic"
	routingKey := payments.CreateBalanceRoutingKey
	body := jsonObj

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req

}
