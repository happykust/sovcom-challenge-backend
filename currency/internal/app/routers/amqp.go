package routers

import (
	APIAmqp "currency/internal/app/api/domain/currency-ws/delivery"
	"currency/internal/app/domain/currency/delivery/amqp"
)

func MainAmqpRouter() {
	go amqp.CurrencyAddTickerConsumer()
	go amqp.CurrencyReadConsumer()
	go amqp.CurrencyDeleteConsumer()
	go amqp.CurrencyCreateGroupConsumer()
	go amqp.CurrencyReadGroupConsumer()
	go amqp.CurrencyDeleteGroupConsumer()
	go amqp.CurrencyReadTickerChangesConsumer()
	go amqp.CurrencyReadTickerConsumer()
	go APIAmqp.CurrencyUpdateConsumer()
	select {}
}
