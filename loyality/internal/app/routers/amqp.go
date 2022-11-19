package routers

import (
	PromocodesAmqp "loyality/internal/app/domain/promocodes/delivery/amqp"
	ReferralAmqp "loyality/internal/app/domain/referral/delivery/amqp"
)

func ReferralAmqpRouter() {
	go ReferralAmqp.ReferralCreateConsumer()
	go ReferralAmqp.ReferralReadConsumer()
	go ReferralAmqp.ReferralAddConsumer()
	go ReferralAmqp.ReferralDeleteConsumer()
	select {}
}

func PromocodesAmqpRouter() {
	go PromocodesAmqp.PromocodesCreateConsumer()
	go PromocodesAmqp.PromocodesReadConsumer()
	go PromocodesAmqp.PromocodesUpdateConsumer()
	go PromocodesAmqp.PromocodesDeleteConsumer()
	go PromocodesAmqp.PromocodesUseConsumer()
	select {}
}
