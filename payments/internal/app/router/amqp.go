package router

import "payments/internal/app/domain/balance/delivery/amqp"

func MainAmqpRouter() {
	go amqp.UserCreatedEventCreateBalance()
	go amqp.GetUserBalance()
	go amqp.UpdateUserBalance()
	select {}
	//go amqp.TestConsumer()
	//select {}
}
