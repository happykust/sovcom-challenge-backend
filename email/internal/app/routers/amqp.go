package routers

import "email/internal/app/domain/smtp/delivery/amqp"

func MainAmqpRouter() {
	go amqp.EmailConsumer()
	go amqp.EmailByUserIDConsumer()
	select {}
}
