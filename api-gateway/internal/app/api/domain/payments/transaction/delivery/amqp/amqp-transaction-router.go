package amqp

import (
	amqp_easier "api-gateway/pkg/core/broker/amqp-easier"
	"libs/contracts/payments"
)

func Create(jsonObj []byte) []byte {
	exchangeName := payments.PaymentsExchange
	exchangeType := "topic"
	routingKey := payments.CreateTransactionTopic
	body := jsonObj
	connName := payments.CreateTransactionConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}

func Get(jsonObj []byte) []byte {
	exchangeName := payments.PaymentsExchange
	exchangeType := "topic"
	routingKey := payments.GetTransactionTopic
	body := jsonObj
	connName := payments.GetTransactionConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}

func Update(jsonObj []byte) []byte {
	exchangeName := payments.PaymentsExchange
	exchangeType := "topic"
	routingKey := payments.UpdateTransactionStatusTopic
	body := jsonObj
	connName := payments.UpdateTransactionStatusConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}
