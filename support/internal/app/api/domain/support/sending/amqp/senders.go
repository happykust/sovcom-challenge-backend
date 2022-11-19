package amqp

import (
	supportToAccounts "libs/contracts/support/accounts"
)

func GetUserDataByAccessToken(request supportToAccounts.ValidateRequest) supportToAccounts.ValidateResponse {
	//TODO: Unmock
	var response supportToAccounts.ValidateResponse
	response = supportToAccounts.ValidateResponse{Status: true, UserID: 123, FirstName: "Bob", LastName: "F",
		Role: "user"}
	return response
}
