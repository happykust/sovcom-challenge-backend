package amqp

import (
	amqp_easier "api-gateway/pkg/core/broker/amqp-easier"
	CurrencyDeals "libs/contracts/deals/currency"
)

func Create(jsonObj []byte) []byte {
	exchangeName := CurrencyDeals.CurrencyDealsExchange
	exchangeType := "topic"
	routingKey := CurrencyDeals.BuyTopic
	body := jsonObj
	connName := CurrencyDeals.BuyConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}

func Delete(jsonObj []byte) []byte {
	exchangeName := CurrencyDeals.CurrencyDealsExchange
	exchangeType := "topic"
	routingKey := CurrencyDeals.DeleteTopic
	body := jsonObj
	connName := CurrencyDeals.DeleteConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}

func Read(jsonObj []byte) []byte {
	exchangeName := CurrencyDeals.CurrencyDealsExchange
	exchangeType := "topic"
	routingKey := CurrencyDeals.ReadTopic
	body := jsonObj
	connName := CurrencyDeals.ReadConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}
