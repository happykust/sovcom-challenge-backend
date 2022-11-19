package router

import "payments/internal/app/domain/balance/delivery/amqp"

func MainAmqpRouter() {
	go amqp.UserCreatedEventCreateBalance()
	select {}
	//go amqp.TestConsumer()
	//select {}
}
