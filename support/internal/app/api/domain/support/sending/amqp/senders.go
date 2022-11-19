package amqp

import (
	supportToAccounts "libs/contracts/support/accounts"
)

func GetUserDataByAccessToken(request supportToAccounts.ValidateRequest) supportToAccounts.ValidateResponse {
	//TODO: Unmock
	var response supportToAccounts.ValidateResponse
	response = supportToAccounts.ValidateResponse{Status: true, UserID: 23, FirstName: "Jo", LastName: "F",
		Role: "admin"}
	return response
}
