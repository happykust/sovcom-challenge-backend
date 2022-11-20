package payments

type TransactionStatus string

const (
	PROCESSING TransactionStatus = "PROCESSING"
	SUCCESS                      = "SUCCESS"
	CANCELED                     = "CANCELED"
)

type CreateTransactionRequest struct {
	UserId    uint    `json:"user_id"`
	Amount    float64 `json:"amount"`
	PromoCode string  `json:"promocode"`
}

type CreateTransactionResponse struct {
	TransactionId     string            `json:"transaction_uuid"`
	UserId            uint              `json:"user_id"`
	Amount            float64           `json:"amount"`
	TransactionStatus TransactionStatus `json:"transaction_status"`
}
