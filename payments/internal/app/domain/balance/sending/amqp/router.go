package amqp

import (
	"fmt"
	amqp_easier "payments/pkg/core/broker/amqp-easier"
)

func TestSendToConst(jsonObj []byte) {
	exchangeName := "account"
	exchangeType := "topic"
	routingKey := "account.login.command"
	body := jsonObj
	connName := "PRECOLLECTOR"

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	fmt.Println(string(req))
}
