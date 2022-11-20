package main

import (
	"payments/internal/app/router"
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
	//s := transactions.TransactionMock(payments.CreateTransactionRequest{UserId: 4, Amount: 1000000000, PromoCode: "dffd"})
	//fmt.Println(s)

	//balance.CreateUserBalance(1)
	//balance.CreateUserWallet(balance.Wallet{BalanceId: 1, Ticker: "ETH", Amount: 100000000000})
	//balance.GetUserBalance(payments.GetBalancesRequest{UserID: 1, Ticker: "ETH"})
	//balance.CreateUserEvent(100)
	//test := balance.UpdateUserBalance(payments.UpdateBalanceRequest{UserID: 4, Ticker: "sdsd", Amount: 10000000000000})
	//fmt.Println(test)
	go router.MainAmqpRouter()
	select {}

}
