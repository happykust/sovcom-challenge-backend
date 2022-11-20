package amqp

import (
	amqp_easier "api-gateway/pkg/core/broker/amqp-easier"
	"libs/contracts/loyality/promocodes"
)

func Create(jsonObj []byte) []byte {
	exchangeName := promocodes.PromocodesExchange
	exchangeType := "topic"
	routingKey := promocodes.CreateTopic
	body := jsonObj
	connName := promocodes.CreateConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}

func Delete(jsonObj []byte) []byte {
	exchangeName := promocodes.PromocodesExchange
	exchangeType := "topic"
	routingKey := promocodes.DeleteTopic
	body := jsonObj
	connName := promocodes.DeleteConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}

func Read(jsonObj []byte) []byte {
	exchangeName := promocodes.PromocodesExchange
	exchangeType := "topic"
	routingKey := promocodes.ReadTopic
	body := jsonObj
	connName := promocodes.ReadConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}

func Update(jsonObj []byte) []byte {
	exchangeName := promocodes.PromocodesExchange
	exchangeType := "topic"
	routingKey := promocodes.UpdateTopic
	body := jsonObj
	connName := promocodes.UpdateConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}

func Use(jsonObj []byte) []byte {
	exchangeName := promocodes.PromocodesExchange
	exchangeType := "topic"
	routingKey := promocodes.UseTopic
	body := jsonObj
	connName := promocodes.UseConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req
}
