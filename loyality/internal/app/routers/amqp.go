package routers

import (
	PromocodesAmqp "loyality/internal/app/domain/promocodes/delivery/amqp"
)

func PromocodesAmqpRouter() {
	go PromocodesAmqp.PromocodesCreateConsumer()
	go PromocodesAmqp.PromocodesReadConsumer()
	go PromocodesAmqp.PromocodesUpdateConsumer()
	go PromocodesAmqp.PromocodesDeleteConsumer()
	go PromocodesAmqp.PromocodesUseConsumer()
	select {}
}
