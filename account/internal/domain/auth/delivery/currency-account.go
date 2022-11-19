package delivery

import (
	"account/internal/domain/auth"
	amqp_easier "account/pkg/core/broker/amqp-easier"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"libs/contracts/account"
	"libs/contracts/currency"
	"libs/contracts/currency/currencyToAccounts"
)

func CurrencyValidateRequest() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(currency.CurrencyToAccountsConsumerName, account.AccountExchange, "topic", currency.CurrencyToAccountsRoutingKey, currency.CurrencyToAccountsQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			validateReq := &currencyToAccounts.ValidateRequest{}
			err := json.Unmarshal(d.Body, validateReq)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(validateReq)
			res := auth.ValidateTokens(validateReq.AccessToken)
			t, err := json.Marshal(&currencyToAccounts.ValidateResponse{Status: res.Status, UserID: res.UserID})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	<-stopChan
}
