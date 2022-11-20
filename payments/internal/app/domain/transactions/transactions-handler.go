package transactions

import (
	"fmt"
	"github.com/google/uuid"
	"libs/contracts/payments"
	"payments/internal/app/domain/balance"
)

func TransactionMock(payload payments.CreateTransactionRequest) payments.CreateTransactionResponse {
	t := CreateTransaction(payload)
	UpdateTransactionStatus(t.TransactionId, payments.SUCCESS)
	s := GetTransaction(t.TransactionId)
	balance.UpdateUserBalance(payments.UpdateBalanceRequest{UserID: s.UserId, Ticker: "RUB", Amount: s.Amount})
	return payments.CreateTransactionResponse{
		TransactionId:     s.TransactionUUID,
		UserId:            s.UserId,
		Amount:            s.Amount,
		TransactionStatus: s.TransactionStatus,
	}

}

func CreateTransaction(payload payments.CreateTransactionRequest) payments.CreateTransactionResponse {
	uuid := uuid.New()
	transaction := Transaction{
		TransactionUUID: uuid.String(),
		UserId:          payload.UserId,
		Amount:          payload.Amount,
		PromoCode:       payload.PromoCode,
	}
	if len(payload.PromoCode) != 0 {
		// TODO: check promo code from promo service
		fmt.Println("promo code is not empty")
	}
	transaction = CreateNewTransaction(transaction)
	fmt.Println(transaction)

	return payments.CreateTransactionResponse{
		TransactionId:     transaction.TransactionUUID,
		UserId:            transaction.UserId,
		Amount:            transaction.Amount,
		TransactionStatus: transaction.TransactionStatus,
	}
}

func UpdateTransactionStatus(Tuuid string, Tstatus payments.TransactionStatus) {
	transaction := GetTransactionRepoByUUID(Tuuid)
	UpdateTransactionStatusRepo(transaction.ID, Tstatus)
	fmt.Println(transaction)
}

func GetTransaction(Tuuid string) Transaction {
	transaction := GetTransactionRepoByUUID(Tuuid)
	return transaction
}
