package balance

import (
	"fmt"
	"libs/contracts/payments"
)

func CreateUserEvent(userId uint) Wallet {
	oldUserBalance := FindUserBalance(userId)
	if len(oldUserBalance) != 0 {
		fmt.Println("User already has a balance")
		return Wallet{}
	}
	balance := CreateUserBalance(userId)
	wallet := CreateUserWallet(Wallet{Ticker: "RUB", Amount: 0, BalanceId: balance.ID})
	return wallet
}

func GetUserBalance(payload payments.GetBalancesRequest) payments.GetBalancesResponse {
	oldUserBalance := FindUserBalance(payload.UserID)
	if len(oldUserBalance) == 0 {
		return payments.GetBalancesResponse{Ticker: payload.Ticker, Amount: 0}
	}
	balance := FindUserBalanceByTicker(payload.UserID, payload.Ticker)
	if len(balance[0].Wallet) == 0 {
		return payments.GetBalancesResponse{Ticker: payload.Ticker, Amount: 0}
	}
	return payments.GetBalancesResponse{Ticker: payload.Ticker, Amount: balance[0].Wallet[0].Amount}
}

func UpdateUserBalance(payload payments.UpdateBalanceRequest) payments.UpdateBalanceResponse {
	userBalance := FindUserBalance(payload.UserID)
	if len(userBalance) == 0 {
		CreateUserEvent(payload.UserID)
		return payments.UpdateBalanceResponse{Status: false, Ticker: payload.Ticker, Amount: 0}
	}
	balance := FindUserBalanceByTicker(payload.UserID, payload.Ticker)
	if len(balance[0].Wallet) == 0 {
		CreateUserWallet(Wallet{Ticker: payload.Ticker, Amount: payload.Amount, BalanceId: userBalance[0].ID})
		return payments.UpdateBalanceResponse{Status: true, Ticker: payload.Ticker, Amount: payload.Amount}
	}
	userWallet := GetUserBalance(payments.GetBalancesRequest{UserID: payload.UserID, Ticker: payload.Ticker})
	newWaller := userWallet.Amount + payload.Amount
	UpdateUserBalanceRep(payload.UserID, payload.Ticker, newWaller)
	return payments.UpdateBalanceResponse{Status: true, Ticker: payload.Ticker, Amount: newWaller}
}
