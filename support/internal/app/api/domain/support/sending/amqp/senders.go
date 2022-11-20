package amqp

import (
	"encoding/json"
	"libs/contracts/account"
	supportToAccounts "libs/contracts/support/accounts"
	amqp_easier "support/pkg/core/broker/amqp-easier"
)

func GetUserDataByAccessToken(request supportToAccounts.ValidateRequest) supportToAccounts.ValidateResponse {
	exchangeName := account.AccountExchange
	exchangeType := "topic"
	routingKey := account.SupportValidateTopic
	body, _ := json.Marshal(request)
	connName := account.SupportValidateConsumerName

	req := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	var response supportToAccounts.ValidateResponse
	err := json.Unmarshal(req, &response)
	if err != nil {
		return supportToAccounts.ValidateResponse{Status: false}
	}
	return response
}
