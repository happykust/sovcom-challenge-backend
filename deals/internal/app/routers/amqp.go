package routers

import (
	currencyDealsAmqp "deals/internal/app/domain/currencyDeals/delivery/amqp"
	simpleDealsAmqp "deals/internal/app/domain/simpleDeals/delivery/amqp"
)

func SimpleDealsAmqpRouter() {
	go simpleDealsAmqp.SimpleDealBuyConsumer()
	select {}
}

func CurrencyDealsAmqpRouter() {
	go currencyDealsAmqp.CurrencyDealBuyConsumer()
	go currencyDealsAmqp.CurrencyDealDeleteConsumer()
	go currencyDealsAmqp.CurrencyDealReadConsumer()
	go currencyDealsAmqp.CurrencyDealCurrencyChangeConsumer()
	select {}
}
