package amqp

import (
	amqp_easier "currency/pkg/core/broker/amqp-easier"
	"encoding/json"
	"libs/contracts/account"
	"libs/contracts/currency"
	"libs/contracts/currency/currencyToAccounts"
)

func GetUserDataByAccessToken(request currencyToAccounts.ValidateRequest) currencyToAccounts.ValidateResponse {
	exchangeName := account.AccountExchange
	exchangeType := "topic"
	routingKey := currency.CurrencyToAccountsRoutingKey
	connName := account.ValidateTokenConsumerName
	body, _ := json.Marshal(request)

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	var response currencyToAccounts.ValidateResponse
	err := json.Unmarshal(req, &response)
	if err != nil {
		return currencyToAccounts.ValidateResponse{Status: false}
	}

	return response
}
