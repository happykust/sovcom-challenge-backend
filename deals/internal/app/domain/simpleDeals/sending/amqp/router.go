package amqp

import (
	amqp_easier "deals/pkg/core/broker/amqp-easier"
	"encoding/json"
	"libs/contracts/payments"
)

func GetUserBalances(request payments.GetBalancesRequest) payments.GetBalancesResponse {
	exchangeName := payments.PaymentsExchange
	exchangeType := "topic"
	routingKey := payments.GetBalanceTopic
	connName := payments.GetBalanceConsumerName
	body, _ := json.Marshal(request)

	var response payments.GetBalancesResponse
	responseAmqp := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)

	err := json.Unmarshal(responseAmqp, &response)
	if err != nil {
		return payments.GetBalancesResponse{}
	}

	return response
}

func UpdateUserBalances(request payments.UpdateBalanceRequest) payments.UpdateBalanceResponse {
	exchangeName := payments.PaymentsExchange
	exchangeType := "topic"
	routingKey := payments.UpdateBalanceTopic
	connName := payments.UpdateBalanceConsumerName
	body, _ := json.Marshal(request)

	var response payments.UpdateBalanceResponse
	responseAmqp := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)

	err := json.Unmarshal(responseAmqp, &response)
	if err != nil {
		return payments.UpdateBalanceResponse{}
	}

	return response
}
