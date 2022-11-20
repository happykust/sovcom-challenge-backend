package amqp

import (
	amqp_easier "api-gateway/pkg/core/broker/amqp-easier"
	"libs/contracts/payments"
)

func Get(jsonObj []byte) []byte {
	exchangeName := payments.PaymentsExchange
	exchangeType := "topic"
	routingKey := payments.GetBalanceTopic
	body := jsonObj
	connName := payments.GetBalanceConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}
