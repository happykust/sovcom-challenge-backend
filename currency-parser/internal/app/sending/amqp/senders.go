package amqp

import (
	amqp_easier "currency-parser/pkg/core/broker/amqp-easier"
	"encoding/json"
	"libs/contracts/currency"
	currencyDeals "libs/contracts/deals/currency"
)

func SendCurrencyUpdateToCurrency(request currency.CurrencyUpdateRequestToCurrency) {
	exchangeName := currency.CurrencyExchange
	exchangeType := "topic"
	routingKey := currency.CurrencyUpdateTopic
	connName := currency.CurrencyUpdateConsumerName
	body, _ := json.Marshal(request)

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	_ = req
}

func SendCurrencyUpdateToDeals(request currency.CurrencyUpdateRequestToDeals) {
	exchangeName := currencyDeals.CurrencyDealsExchange
	exchangeType := "topic"
	routingKey := currencyDeals.CurrencyChangeTopic
	connName := currencyDeals.CurrencyChangeConsumerName
	body, _ := json.Marshal(request)

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	_ = req
}
