package amqp

import "libs/contracts/currency/currencyToAccounts"

func GetUserDataByAccessToken(request currencyToAccounts.ValidateRequest) currencyToAccounts.ValidateResponse {
	//TODO: Unmock
	var response currencyToAccounts.ValidateResponse
	response = currencyToAccounts.ValidateResponse{Status: true, UserID: 123}
	return response
}