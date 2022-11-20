package router

import (
	balanceAmqp "payments/internal/app/domain/balance/delivery/amqp"
	transactionsAmqp "payments/internal/app/domain/transactions/delivery/amqp"
)

func MainAmqpRouter() {
	go balanceAmqp.UserCreatedEventCreateBalance()
	go balanceAmqp.GetUserBalance()
	go balanceAmqp.UpdateUserBalance()
	go transactionsAmqp.CreateConsumer()
	go transactionsAmqp.UpdateTransactionStatusConsumer()
	go transactionsAmqp.GetConsumer()
	select {}
}
