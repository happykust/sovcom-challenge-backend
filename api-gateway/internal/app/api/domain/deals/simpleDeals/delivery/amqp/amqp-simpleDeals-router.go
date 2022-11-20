package amqp

import (
	amqp_easier "api-gateway/pkg/core/broker/amqp-easier"
	SimpleDeals "libs/contracts/deals/simple"
)

func Create(jsonObj []byte) []byte {
	exchangeName := SimpleDeals.SimpleDealsExchange
	exchangeType := "topic"
	routingKey := SimpleDeals.BuyTopic
	body := jsonObj
	connName := SimpleDeals.BuyConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}
