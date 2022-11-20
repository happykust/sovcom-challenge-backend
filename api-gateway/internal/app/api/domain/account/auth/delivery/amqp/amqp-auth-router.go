package amqp

import (
	amqp_easier "api-gateway/pkg/core/broker/amqp-easier"
	"fmt"
	"libs/contracts/account"
)

func Register(jsonObj []byte) []byte {
	exchangeName := account.AccountExchange
	exchangeType := "topic"
	routingKey := account.SignUpTopic
	body := jsonObj
	connName := account.SignUpConsumerName
	fmt.Println("Register")

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)

	return req
}

func Login(jsonObj []byte) []byte {
	exchangeName := account.AccountExchange
	exchangeType := "topic"
	routingKey := account.SignIn
	body := jsonObj
	connName := account.SignInConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req

}

func Refresh(jsonObj []byte) []byte {
	exchangeName := account.AccountExchange
	exchangeType := "topic"
	routingKey := account.RefreshTopic
	body := jsonObj
	connName := account.RefreshConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req

}

func Logout(jsonObj []byte) []byte {
	exchangeName := account.AccountExchange
	exchangeType := "topic"
	routingKey := account.LogoutTopic
	body := jsonObj
	connName := account.LogoutConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	return req

}
