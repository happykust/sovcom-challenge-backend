package transactions

import (
	"libs/contracts/payments"
	"payments/pkg/core/database"
)

func CreateNewTransaction(payload Transaction) Transaction {
	database.PG.Create(&payload)
	return payload
}

func UpdateTransactionStatusRepo(TransactionId uint, status payments.TransactionStatus) {
	var transaction Transaction
	database.PG.Where("id = ?", TransactionId).Find(&transaction)
	transaction.TransactionStatus = status
	database.PG.Save(&transaction)

}

func GetTransactionRepoByUser(userId uint) []Transaction {
	var transaction []Transaction
	database.PG.Where("user_id = ?", userId).Find(&transaction)
	return transaction
}

func GetTransactionRepoByUUID(TransactionId string) Transaction {
	var transaction Transaction
	database.PG.Where("transaction_uuid = ?", TransactionId).Find(&transaction)
	return transaction
}
