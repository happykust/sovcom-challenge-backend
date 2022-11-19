package amqp

import (
	amqp_easier "currency-parser/pkg/core/broker/amqp-easier"
	"encoding/json"
	"libs/contracts/currency"
)

func SendCurrencyUpdate(request currency.CurrencyUpdateRequest) {
	exchangeName := currency.CurrencyExchange
	exchangeType := "topic"
	routingKey := currency.CurrencyUpdateTopic
	connName := currency.CurrencyUpdateConsumerName
	body, _ := json.Marshal(request)

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	_ = req
}
