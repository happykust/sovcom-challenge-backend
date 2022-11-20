package router

import "account/internal/domain/auth/delivery"

func AmqpMainRouter() {
	go delivery.CurrencyValidateRequest()
	select {}
}
