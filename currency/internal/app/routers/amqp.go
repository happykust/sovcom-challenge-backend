package routers

import (
	APIAmqp "currency/internal/app/api/domain/currency-ws/delivery"
)

func MainAmqpRouter() {
	go APIAmqp.CurrencyUpdateConsumer()
	select {}
}
