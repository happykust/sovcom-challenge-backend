package payments

var PaymentsExchange = "payments"

// Get balance

var CreateBalanceQueue = "payments.service.create.balance.queue"
var CreateBalanceRoutingKey = "payments.service.create.balance.key"
var CreateBalanceQueueName = "payments.service.create.balance.queuename"

var GetBalanceTopic = "command.payments.balance.get"

var GetBalanceConsumerName = "payments.service.consumer.balance.get"

var GetBalanceQueueName = "payments.service.queue.balance.get"

// UpdateBalance

var UpdateBalanceTopic = "command.payments.balance.update"

var UpdateBalanceConsumerName = "payments.service.consumer.balance.update"

var UpdateBalanceQueueName = "payments.service.queue.balance.update"

///

var CreateTransactionTopic = "command.payments.transactions.update"

var CreateTransactionConsumerName = "payments.service.consumer.transactions.update"

var CreateTransactionQueueName = "payments.service.queue.transactions.update"

var UpdateTransactionStatusTopic = "command.payments.transactions.update"

var UpdateTransactionStatusName = "payments.service.consumer.transactions.update"

var UpdateTransactionStatusQueueName = "payments.service.queue.transactions.update"

var UpdateTransactionStatusConsumerName = "payments.service.consumer.transactions.update"

var GetTransactionTopic = "command.payments.balance.update"

var GetTransactionConsumerName = "payments.service.consumer.transactions.update"

var GetTransactionQueueName = "payments.service.queue.transactions.update"
