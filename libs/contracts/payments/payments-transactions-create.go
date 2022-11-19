package payments

type TransactionStatus string

const (
	PROCESSING TransactionStatus = "PROCESSING"
	SUCCESS                      = "SUCCESS"
	CANCELED                     = "CANCELED"
)

type CreateTransactionRequest struct {
	UserId    uint
	Amount    float64
	PromoCode string
}

type CreateTransactionResponse struct {
	TransactionId     string
	UserId            uint
	Amount            float64
	TransactionStatus TransactionStatus
}
