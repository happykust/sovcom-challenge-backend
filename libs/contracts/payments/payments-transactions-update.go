package payments

type PaymentsTransactionsUpdateRequest struct {
	TransactionUUID   string `json:"transaction_uuid"`
	TransactionStatus string `json:"transaction_status"`
}

type PaymentsTransactionsUpdateResponse struct {
	Message string `json:"message"`
}
