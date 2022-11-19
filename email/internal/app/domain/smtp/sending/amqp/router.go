package amqp

import (
	emailAccounts "libs/contracts/email/accounts"
)

// TODO: Create consumer in accounts to get email from user id.

func GetEmailByUserIDFromAccounts(request emailAccounts.SearchEmailRequest) emailAccounts.SearchEmailResponse {
	//exchangeName := email.EmailExchange
	//exchangeType := "topic"
	//routingKey := email.GetEmailTopic
	//connName := email.GetEmailConsumerName
	//body, _ := json.Marshal(request)

	var response emailAccounts.SearchEmailResponse
	//responseAmqp := amqp_easier.PublishConstructor(connName, exchangeName, exchangeType, &routingKey, &body)
	//
	//err := json.Unmarshal(responseAmqp, &response)
	//if err != nil {
	//	return emailAccounts.SearchEmailResponse{}
	//}

	return response
}
