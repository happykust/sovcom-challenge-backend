package routers

import (
	currencyDealsAmqp "deals/internal/app/domain/currencyDeals/delivery/amqp"
	simpleDealsAmqp "deals/internal/app/domain/simpleDeals/delivery/amqp"
)

func SimpleDealsAmqpRouter() {
	go simpleDealsAmqp.SimpleDealBuyConsumer()
	go simpleDealsAmqp.SimpleDealSellConsumer()
	select {}
}

func CurrencyDealsAmqpRouter() {
	go currencyDealsAmqp.CurrencyDealBuyConsumer()
	go currencyDealsAmqp.CurrencyDealSellConsumer()
	go currencyDealsAmqp.CurrencyDealDeleteConsumer()
	go currencyDealsAmqp.CurrencyDealReadConsumer()
	go currencyDealsAmqp.CurrencyDealCurrencyChangeConsumer()
	select {}
}
