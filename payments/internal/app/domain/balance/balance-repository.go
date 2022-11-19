package balance

import "payments/pkg/core/database"

func CreateUserBalance(userId uint) Balance {
	var userBalance Balance
	userBalance.UserId = userId
	database.PG.Create(&userBalance)
	return userBalance
}

func CreateUserWallet(wallet Wallet) Wallet {
	database.PG.Create(&wallet)
	return wallet
}

func FindUserBalanceByTicker(userId uint, ticker string) []Balance {
	var userBalance []Balance
	database.PG.Model(&userBalance).Preload("Wallet", "ticker = ?", ticker).Where("user_id = ?", userId).Find(&userBalance)
	return userBalance
}

func UpdateUserBalanceRep(userId uint, ticker string, amount float64) Balance {
	var userBalance Balance
	database.PG.Exec("UPDATE wallets SET amount = ? WHERE balance_id = ? AND ticker = ?", amount, userId, ticker)
	return userBalance
}
func FindUserBalance(userId uint) []Balance {
	var userBalance []Balance
	database.PG.Where("user_id = ?", userId).Find(&userBalance)
	return userBalance
}
