package main

import (
	"fmt"
	"libs/contracts/payments"
	"payments/internal/app/domain/transactions"
	"payments/pkg/core/config"
	"payments/pkg/core/database"
	"payments/pkg/core/database/migrations"
)

func init() {
	config.Init()
}

func main() {
	database.OpenPostgresConnect()
	migrations.InitMigrations()
	s := transactions.TransactionMock(payments.CreateTransactionRequest{UserId: 4, Amount: 1000000000, PromoCode: "ПИДОРАЗЗЗЗЗЗ"})
	fmt.Println(s)

	//balance.CreateUserBalance(4)
	//balance.CreateUserWallet(balance.Wallet{BalanceId: 4, Ticker: "USD", Amount: 1000})
	//balance.GetUserBalance(payments.PaymentsCheckBalanceByRequest{UserID: 4, Currency: "RUB"})
	//balance.GetUserBalance(payments.PaymentsCheckBalanceByRequest{UserID: 4, Currency: "USD"})
	//balance.CreateUserEvent(100)
	//test := balance.UpdateUserBalance(payments.UpdateBalanceRequest{UserID: 4, Ticker: "PENICCOIN", Amount: 10000000000000})
	//fmt.Println(test)

}